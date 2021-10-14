using System;
using System.Collections.Concurrent;
using System.Linq;
using System.Net.Http;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;

namespace Hackube.ServiceTest
{
    class Program
    {
        static async Task Main(string[] args)
        {
            if (args.Length == 0)
            {
                args = new[] {"http://localhost:32005/env/hostname"};
                // Console.WriteLine("Testing Url Not Provided");
            }

            var testConfig = new TestConfig()
            {
                Uri = args[0]
            };

            var builder = new HostBuilder();
            builder.ConfigureLogging((builder) =>
            {
                builder.AddConsole();
                builder.SetMinimumLevel(LogLevel.Information);
            }).ConfigureServices(((context, collection) =>
            {
                collection.AddSingleton(testConfig);
                collection.AddHostedService<LoadBalanceTestService>();
            }));

            Console.WriteLine("Starting Testing");
            await builder.RunConsoleAsync();
        }
    }

    public class TestConfig
    {
        public string Uri { get; set; }
    }

    public class LoadBalanceTestService : IHostedService
    {
        private readonly ILogger<LoadBalanceTestService> _logger;
        private readonly TestConfig _config;

        public LoadBalanceTestService(ILogger<LoadBalanceTestService> logger, TestConfig config)
        {
            _logger = logger;
            _config = config;
        }

        private ConcurrentDictionary<string, int> _grpHostnames = new ConcurrentDictionary<string, int>();

        public Task Run()
        {
            _logger.LogInformation("Starting LoadBalance Testing");
            HttpClient client = new HttpClient();
            ParallelEnumerable.Range(1, 1000)
                .WithMergeOptions(ParallelMergeOptions.FullyBuffered)
                .ForAll(async x => await CollectCounter(x, client).ConfigureAwait(false));
            //.ForAll(i => CollectCounter(client).ConfigureAwait(false));

            while (_grpHostnames.Values.Sum() != 1000)
            {
            }

            foreach (var hostName in _grpHostnames)
            {
                _logger.LogInformation($"{hostName.Key} -> {hostName.Value}");
            }

            return Task.CompletedTask;
        }

        public async Task CollectCounter(int index, HttpClient client)
        {
            var content = await client.GetStringAsync(_config.Uri);
            _grpHostnames.AddOrUpdate(content, (key) => 1, (key, curVal) => curVal + 1);
            // _logger.LogInformation($"{index} : {content}");
        }

        public async Task StartAsync(CancellationToken cancellationToken)
        {
            await Run();
        }

        public Task StopAsync(CancellationToken cancellationToken)
        {
            return Task.CompletedTask;
        }
    }
}