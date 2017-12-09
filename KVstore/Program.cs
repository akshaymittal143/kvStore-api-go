﻿using Microsoft.Owin.Hosting;
using System;
using System.Net.Http;

namespace KVstore
{
    public class Program
    {
        static void Main(string[] args)
        {
            string baseAddress1 = "http://localhost:9001/";
            string baseAddress2 = "http://localhost:9002/";
            string baseAddress3 = "http://localhost:9003/";
            string baseAddress4 = "http://localhost:9004/";
            string baseAddress5 = "http://localhost:9005/";

            // Start OWIN host1 
            using (WebApp.Start<Startup>(url: baseAddress1))
            {
                // Create HttpCient and make a request to api/values 
                HttpClient client1 = new HttpClient();

                var response1 = client1.GetAsync(baseAddress1 + "api/values").Result;

                Console.WriteLine(response1);
                Console.WriteLine(response1.Content.ReadAsStringAsync().Result);
                Console.ReadLine();
            }

            // Start OWIN host2 
            using (WebApp.Start<Startup1>(url: baseAddress2))
            {
                // Create HttpCient and make a request to api/values 
                HttpClient client2 = new HttpClient();

                var response2 = client2.GetAsync(baseAddress2 + "api/values").Result;

                Console.WriteLine(response2);
                Console.WriteLine(response2.Content.ReadAsStringAsync().Result);
                Console.ReadLine();
            }
            // Start OWIN host3 
            using (WebApp.Start<Startup>(url: baseAddress3))
            {
                // Create HttpCient and make a request to api/values 
                HttpClient client3 = new HttpClient();

                var response3 = client3.GetAsync(baseAddress3 + "api/values").Result;

                Console.WriteLine(response3);
                Console.WriteLine(response3.Content.ReadAsStringAsync().Result);
                Console.ReadLine();
            }
            // Start OWIN host4 
            using (WebApp.Start<Startup>(url: baseAddress4))
            {
                // Create HttpCient and make a request to api/values 
                HttpClient client4 = new HttpClient();

                var response4 = client4.GetAsync(baseAddress4 + "api/values").Result;

                Console.WriteLine(response4);
                Console.WriteLine(response4.Content.ReadAsStringAsync().Result);
                Console.ReadLine();
            }
            // Start OWIN host5 
            using (WebApp.Start<Startup>(url: baseAddress5))
            {
                // Create HttpCient and make a request to api/values 
                HttpClient client5 = new HttpClient();

                var response5 = client5.GetAsync(baseAddress5 + "api/values").Result;

                Console.WriteLine(response5);
                Console.WriteLine(response5.Content.ReadAsStringAsync().Result);
                Console.ReadLine();
            }


        }
    }
}
