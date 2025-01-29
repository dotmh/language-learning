namespace CurrencyCalculator
{
    public class Denomination
    {
        public required string Name { get; set; }
        public int Value { get; set; }
    }

    public class Denominations
    {
        public required List<Denomination> Notes { get; set; }
        public required List<Denomination> Coins { get; set; }
    }

    public class WalletItem
    {
        public int Value { get; set; }
        public int Qty { get; set; }
    }

    public class Wallet
    {
        public required List<WalletItem> Notes { get; set; }
        public required List<WalletItem> Coins { get; set; }
    }

    public class CurrencyAPIResponse
    {
        public required Denominations Denominations { get; set; }
        public required Wallet Wallet { get; set; }
    }

    public class FreeCurrencyCurrency
    {
        public required decimal GBP { get; set; }
    }

    public class FreeCurrencyAPIResponse
    {
        public required FreeCurrencyCurrency Data { get; set; }
    }
}