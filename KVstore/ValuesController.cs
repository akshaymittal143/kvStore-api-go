using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Web.Http;

namespace KVstore
{
    public class ValuesController : ApiController
    {
        static Dictionary<int, string> dictionary = new Dictionary<int, string>();

        // GET api/values 
        public Dictionary<int, string> Get()
        {
            //return new string[] { "value1", "value2" };
            return dictionary;
        }

        // GET api/values/5 
        public string Get(int id)
        {

            string result;
            if (dictionary.TryGetValue(id, out result))
            {
                return result;
            }

            HttpResponseMessage response = Request.CreateResponse(HttpStatusCode.NotFound, "Not Found");
            return response.ToString();
        }

        // POST api/values 
        public string Post([FromBody]string value)
        {
            int count = dictionary.Count;
            dictionary.Add(count + 1, value);
            HttpResponseMessage Sucessresponse = Request.CreateResponse(HttpStatusCode.Created, "Created Sucessfully");
            return Sucessresponse.ToString();
        }

        // PUT api/values/5 
        public string Put(int id, [FromBody]string value)
        {
            HttpResponseMessage Sucessresponse = Request.CreateResponse(HttpStatusCode.OK, "Updated Sucessfully");
            dictionary[id] = value;
            return Sucessresponse.ToString();
        }

        // DELETE api/values/5 
        public string Delete(int id)
        {
            if (dictionary.ContainsKey(id))
            {
                dictionary.Remove(id);
                HttpResponseMessage Sucessresponse = Request.CreateResponse(HttpStatusCode.OK, "Deleted Sucessfully");
                return Sucessresponse.ToString();
            }
            return HttpStatusCode.NotFound.ToString();
        }
    }
}
