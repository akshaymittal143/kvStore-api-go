using Microsoft.Owin.Hosting;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading;

namespace KVstore
{
    public class Program
    {
        private static readonly ManualResetEvent _wait = new ManualResetEvent(initialState: false);

        static void Main(string[] args)
        {
            Console.CancelKeyPress += Console_CancelKeyPress;

            var baseAddresses = new[]
            {
                "http://localhost:9001/",
                "http://localhost:9002/",
                "http://localhost:9003/",
                "http://localhost:9004/",
                "http://localhost:9005/"
            };

            var httpListeners = new List<IDisposable>();
            foreach (var baseAddress in baseAddresses)
            {
                var listener = WebApp.Start<Startup>(baseAddress);
                httpListeners.Add(listener);
            }

            foreach (var baseAddress in baseAddresses)
            {
                HttpClient client1 = new HttpClient();

                var response1 = client1.GetAsync(baseAddress + "api/values").Result;

                Console.WriteLine(response1);
                Console.WriteLine(response1.Content.ReadAsStringAsync().Result);
            }

            Console.WriteLine("Listening... Press ^C to quit");

            _wait.WaitOne();

            Console.WriteLine("Shutting down...");

            for (int i = 0; i < httpListeners.Count; i++)
            {
                Console.WriteLine("Stopping listener {0}...", i + 1);

                var listener = httpListeners[i];
                listener.Dispose();
            }

            Console.WriteLine("Shutdown complete.");
        }

        private static void Console_CancelKeyPress(object sender, ConsoleCancelEventArgs e)
        {
            e.Cancel = true;
            _wait.Set();
        }
    }
}