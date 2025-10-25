# Marketplace API
Marketplace API yang menghubungkan merchant dan customer dengan fitur diskon dan bebas ongkir. Dibangun dengan Golang, Fiber Framework, dan PostgreSQL.

🚀 Fitur
Authentication & Authorization
✅ JWT Token Authentication

✅ Role-based Access Control (Merchant & Customer)

✅ Secure API Endpoints

Merchant Features
✅ Create, Read, Update, Delete Products

✅ View All Transactions

✅ View Sales Reports

✅ See Customers Who Purchased

✅ Top Products Analytics

Customer Features
✅ Browse Products

✅ Purchase Products

✅ Transaction History

✅ Automatic Promotions

Promotion System
✅ Free Shipping: Transaksi produk di atas Rp 15,000

✅ 10% Discount: Transaksi produk di atas Rp 50,000

🛠️ Teknologi
Backend: Golang 1.21+

Framework: Fiber v2

Database: PostgreSQL

Authentication: JWT

Password Hashing: bcrypt

📋 Prerequisites
Go 1.21 atau lebih tinggi

PostgreSQL

Git

⚙️ Installation
Clone Repository

bash
git clone https://github.com/stefsav-dev/marketplace-api.git
cd marketplace-api
Install Dependencies

bash
go mod tidy
Setup Database

sql
CREATE DATABASE marketplace;
Environment Variables (buat file .env)

env
DATABASE_URL=host=localhost user=postgres password=yourpassword dbname=marketplace port=5432 sslmode=disable TimeZone=Asia/Jakarta
JWT_SECRET=your-super-secret-key-here
Run Application

bash
go run main.go
Server akan berjalan di http://localhost:3000

🎯 Promotion Rules
Free Shipping
Berlaku ketika total transaksi produk > Rp 15,000

Biaya pengiriman menjadi Rp 0

10% Discount
Berlaku ketika total transaksi produk > Rp 50,000

Mendapatkan diskon 10% dari total harga produk

Bisa dikombinasikan dengan free shipping

Contoh Perhitungan:
Transaksi: Rp 40,000

Free Shipping: ✅ (40,000 > 15,000)

Discount: ❌ (40,000 < 50,000)

Final Price: Rp 40,000

Transaksi: Rp 60,000

Free Shipping: ✅ (60,000 > 15,000)

Discount: ✅ (60,000 > 50,000) = Rp 6,000

Final Price: Rp 54,000


Test Scenarios:
✅ Transaksi < 15,000 (No promo)

✅ Transaksi > 15,000 (Free shipping only)

✅ Transaksi > 50,000 (Free shipping + 10% discount)