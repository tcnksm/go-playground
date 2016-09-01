# Certificate

Public Key signed by CA (Certification authority)

## Why ?

To reduce, risk of man-in-the-middle attack.

## How works ?

> Alice wants to encrypt message by Bob's public key and send it to bob. Then bob decrypt it by his private key

1. CA create pub & priv key
1. Bob create pub & priv key
1. Bob register pub key at CA
1. CA sign Bob's pub key by CA's priv key and create certificate
1. Alice gets Bob's certificate(pub key)
1. Alice verifies certificate(pub key) by CA's pub key
1. Alice encrypts message by Bob's pub key
1. Alice sends encrypted message to Bob
1. Bob decrypt message by Bob's priv key

## How to veryfy Bob's cert

How to cert is signed.

1. Bob's cert(pub key) is signed by Sapporo CA's priv key
1. Sapporo CA's cert(pub key) is signed by Hokkaido CA priv key
1. Hokkaido CA's cert(pub key) is signed by Japan CA(Root CA)'s priv key
1. Japan CA(Root CA)'s cert(pub key) is signed by self priv key

How to veryfy Bob's cert signature.

1. Alice gets Bob's cert (pub key) signed by Sapporo CA
1. Alice gets Sapporo CA cert (pub key) signed by Hokkaido's CA
1. Alice gets Hokkaido CA cert (pub key) signed by Japan CA
1. Alice verifies Hokkaido CA cert (pub key) by Japan CA pub key
1. Alice verifies Sapporo CA cert (pub key) by Hokkaido CA pub key
1. Alice verifies Bob's cert by Sapporo CA pub key

## x509

x509の証明書はASN.1（Abstract Syntax Notation One）で表記される．ASN.1 は情報の抽象構文を定義するが情報のEncodeのフォーマットは限定しない．x509ではDER（Distinguished Encoding Rules）でエンコーディングが行われる．Goでは`encoding/ans1`にDERのエンコーダーが準備されている．またASN.1のlow levelの構造のパーサーが`crypto/x509/pkix`に定義されている．証明書や鍵はPEM（Privacy Enhanced Mail）形式でエンコーディングされてファイルに保存される．これらは`encoding/pem`に定義されている．
