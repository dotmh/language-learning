using Newtonsoft.Json;

namespace CurrencyCalculator
{
    public class JSONRequest<T>
    {
        private static readonly HttpClient client = new HttpClient();
        public async Task<T> GetData(string url)
        {
            HttpResponseMessage response = await client.GetAsync(url);
            response.EnsureSuccessStatusCode();
            string responseBody = await response.Content.ReadAsStringAsync();

            if (responseBody == null)
            {
                throw new Exception("No data found");
            }

            T? data = JsonConvert.DeserializeObject<T>(responseBody);

            if (data == null)
            {
                throw new Exception("No data found");
            }

            return data;
        }
    }
}