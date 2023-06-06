import { serve } from "https://deno.land/std@0.155.0/http/server.ts";

interface Denomination {
    name: string,
    value: number
}

interface Denominations {
    notes: Denomination[],
    coins: Denomination[]
}

interface WalletItem {
    value: number,
    qty: number
}

interface Wallet {
    notes: WalletItem[],
    coins: WalletItem[]
}

interface APIResponse {
    denominations: Denominations,
    wallet: Wallet
}

type DenominationFactory = (name: string, value: number) => Denomination;

const denomination: DenominationFactory = (name, value) => ({
    name, value
});

const item = (value: number, qty: number): WalletItem => ({
    value, qty
})

const coin: DenominationFactory = (name, value) => denomination(`${name} coin`, value);
const note: DenominationFactory = (name, value) => denomination(`${name} note`, value);

const yenDenominations: Denominations = {
    notes: [
        note("¥10,000", 10000),
        note("¥5000", 5000),
        note("¥1000", 1000)
    ],
    coins: [
        coin("¥500", 500),
        coin("¥100", 100),
        coin("¥50", 50),
        coin("¥10", 10),
        coin("¥5", 5),
        coin("¥1", 1)
    ]
};

const wallet: Wallet = {
    notes: [item(10000, 5), item(5000, 1), item(1000, 5)],
    coins: [item(500, 1), item(100, 4), item(50, 5), item(10, 8), item(5, 7), item(1, 14)]
};

serve(() => {
    const data: APIResponse = {
        denominations: yenDenominations,
        wallet
    }
    return new Response(
        JSON.stringify(data),
        { headers: { "content-type": "application/json" } }
    );
});