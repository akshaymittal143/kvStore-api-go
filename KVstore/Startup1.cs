using Owin;
using System.Web.Http;

namespace KVstore
{
    class Startup1
    {
        public void Configuration(IAppBuilder appBuilder)
        {
            // Configure Web API for self-host. 
            HttpConfiguration config2 = new HttpConfiguration();
            config2.Routes.MapHttpRoute(
                name: "DefaultApi",
                routeTemplate: "api/{controller}/{id}",
                defaults: new { id = RouteParameter.Optional }
            );

            appBuilder.UseWebApi(config2);
        }
    }
}
