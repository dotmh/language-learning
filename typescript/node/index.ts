(async () => {
    const { FREE_CURRENCY_API_KEY } = process.env;
    const url = `https://api.freecurrencyapi.com/v1/latest?apikey=${FREE_CURRENCY_API_KEY}&currencies=GBP&base_currency=JPY`

    const response = await fetch(url);
    const { data } = await response.json();
    const { GBP } = data;

    const Y_1 = +GBP;

    const N_10000 = 10000 * 5;
    const N_1000 = 1000 * 5;
    const N_5000 = 5000 * 1;

    const C_500 = 500 * 1;
    const C_100 = 100 * 4;
    const C_50 = 50 * 5;
    const C_10 = 10 * 8;
    const C_5 = 5 * 7;
    const C_1 = 1 * 14;

    const parts = [
        N_10000,
        N_5000,
        N_1000,
        C_500,
        C_100,
        C_50,
        C_10,
        C_5,
        C_1
    ];

    console.log("Wallet", parts.map((t) => `¥${t}`).join(' | '))

    const total = parts.reduce((a, i) => a + i);
    const totalInGBP = total * Y_1;
    const totalInGBPto2dec = Math.round((totalInGBP + Number.EPSILON) * 100) / 100;

    console.log(`You have ¥${total}, which is £${totalInGBPto2dec}`);
})();