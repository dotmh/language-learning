using System.Globalization;
using Newtonsoft.Json;

namespace CurrencyCalculator
{

    public class Currency
    {
        private static readonly HttpClient client = new HttpClient();

        public async Task processWallet()
        {
            JSONRequest<CurrencyAPIResponse> request = new JSONRequest<CurrencyAPIResponse>();
            CurrencyAPIResponse data = await request.GetData("https://learn.api.dotmh.dev");

            string APIKey = GetFreeCurrencyApiKey();
            string FreeCurrencyAPIURL = $"https://api.freecurrencyapi.com/v1/latest?apikey={APIKey}&currencies=GBP&base_currency=JPY";

            JSONRequest<FreeCurrencyAPIResponse> freeCurrencyRequest = new JSONRequest<FreeCurrencyAPIResponse>();
            FreeCurrencyAPIResponse freeCurrencyResponse = await freeCurrencyRequest.GetData(FreeCurrencyAPIURL);

            int NotesTotal = data.Wallet.Notes.Sum(x => x.Value * x.Qty);
            int CoinsTotal = data.Wallet.Coins.Sum(x => x.Value * x.Qty);
            int WalletTotal = NotesTotal + CoinsTotal;
            decimal Total = WalletTotal * freeCurrencyResponse.Data.GBP;

            string SWalletTotal = WalletTotal.ToString("C", CultureInfo.GetCultureInfoByIetfLanguageTag("ja-JP"));
            string STotal = Total.ToString("C", CultureInfo.GetCultureInfoByIetfLanguageTag("en-GB"));

            Console.WriteLine($"You have {SWalletTotal} which is {STotal}");
        }

        private string GetFreeCurrencyApiKey()
        {
            string? Key = Environment.GetEnvironmentVariable("FREE_CURRENCY_API_KEY");
            if (string.IsNullOrEmpty(Key))
            {
                throw new Exception("No API key found");
            }
            return Key;
        }
    }

}