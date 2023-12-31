# Backend-Marketplace

## Description

Backend-Marketplace adalah bagian dari sistem yang mengelola operasi marketplace. Ini bertindak sebagai backend untuk pengelolaan toko, produk, keranjang belanja, transaksi, ulasan produk, serta wilayah seperti provinsi dan kota. Backend ini berfungsi sebagai inti dari operasi e-commerce, memungkinkan pengguna untuk membeli, menjual, dan melacak produk dan transaksi mereka.

## Feature

### Get Province List

Mengambil daftar provinsi yang tersedia di sistem.

#### Endpoint

```http
GET http://localhost:5000/api/region/province
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": [
		{
			"id": "1",
			"province": "Bali"
		},
		{
			"id": "2",
			"province": "Bangka Belitung"
		},
		{
			"id": "3",
			"province": "Banten"
		}
		// ... (data provinsi lainnya)
	]
}
```

##

### Get City By Province Id

Mengambil daftar kota/kabupaten berdasarkan ID provinsi yang diberikan.

#### Endpoint

```http
GET http://localhost:5000/api/region/city/6
```

#### Parameters

- **6** (number,required): ID provinsi yang digunakan untuk mengambil daftar kota/kabupaten yang berada dalam provinsi tersebut.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": [
		{
			"id": "151",
			"type": "Kota",
			"city_name": "Jakarta Barat",
			"postal_code": "11220"
		},
		{
			"id": "152",
			"type": "Kota",
			"city_name": "Jakarta Pusat",
			"postal_code": "10540"
		},
		{
			"id": "153",
			"type": "Kota",
			"city_name": "Jakarta Selatan",
			"postal_code": "12230"
		}
		// ... (data city lainnya)
	]
}
```

##

### Register User

Digunakan untuk mendaftarkan pengguna baru.

#### Endpoint

```http
POST http://localhost:5000/api/user/register
```

#### Request Body

- **username** (string, required): Nama pengguna.
- **email** (string, required): Alamat email pengguna.
- **password** (string, required): Kata sandi pengguna.
- **city_id** (string,required): id city pengguna
- **phone_number** (string, optional): No telephone pengguna.
- **address** (string, optional): Alamat pengguna (dapat dikosongkan).
- **image_url** (string, optional): URL gambar profil pengguna (dapat dikosongkan).

#### Contoh Request Body:

```json
{
	"username": "person",
	"email": "person@gmail.com",
	"password": "123",
	"city_id": "153",
	"phone_number": "",
	"address": "",
	"image_url": ""
}
```

#### Response

- **HTTP Status**: 201 Created
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success Registration User"
}
```

> [!NOTE]
> Kata sandi (password) yang digunakan dalam contoh di atas telah di-hash dengan algoritma bcrypt.
> jika user memasukkan email dan no telephone yang sudah ada maka akan mengembalikan pesan error.

##

### Login User

Digunakan untuk mengautentikasi pengguna.

#### Endpoint

```http
POST http://localhost:5000/api/user/login
```

#### Request Body

- **email** (string, required): Alamat email pengguna.
- **password** (string, required): Kata sandi pengguna.

#### Contoh Request Body:

```json
{
	"email": "person@gmail.com",
	"password": "123"
}
```

### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8
- **Set-Cookie**: tkn_ck=[token]; Path=/; Domain=localhost; Max-Age=86400; HttpOnly

```json
{
	"your_access_token": "[token]"
}
```

> [!NOTE]
> Token yang diset dicookie dan token yang dikembalikan ke dalam response memiliki value yang berbeda

##

### Get User Access Token

Digunakan untuk mendapatkan akses token pengguna yang sah.

#### Endpoint

```http
GET http://localhost:5000/api/user/token
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"your_access_token": "[token]"
}
```

##

### Logout User

Digunakan untuk logout (keluar) pengguna dari aplikasi.

#### Endpoint

```http
DELETE http://localhost:5000/api/user/logout
```

### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8
- **Set-Cookie**: tkn_ck=; Path=/; Domain=localhost; Max-Age=0; HttpOnly

```json
{
	"msg": "Successfull logout"
}
```

##

### Update User

Digunakan untuk memperbarui informasi pengguna yang ada.

#### Endpoint

```http
PUT http://localhost:5000/api/user/adf4b794-f398-464a-b6d5-ef8a078f0705
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **adf4b794-f398-464a-b6d5-ef8a078f0705** (string,required): User ID.

#### Request Body

- **username** (string, optional): Nama pengguna (dapat dikosongkan).
- **email** (string, optional): Alamat email pengguna (dapat dikosongkan).
- **password** (string, optional): Kata sandi pengguna (dapat dikosongkan).
- **city_id** (string,optional): city id pengguna (dapat dikosongkan)
- **phone_number** (string, optional): Nomor telepon pengguna (dapat dikosongkan).
- **address** (string, optional): Alamat pengguna (dapat dikosongkan).
- **image_url** (string, optional): URL gambar profil pengguna (dapat dikosongkan).

#### Contoh Request Body:

```json
{
	"username": "",
	"email": "person@gmail.com",
	"password": "",
	"city_id": "153",
	"phone_number": "",
	"address": "Jl. Buntu",
	"image_url": ""
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "Successfull Update User With Id 'adf4b794-f398-464a-b6d5-ef8a078f0705'",
	"data": {
		"username": "person",
		"email": "person@gmail.com",
		"password": "$2a$10$lQk04EeOxkf8HR1IqoZBXuhEHThgc6OT2SmVR8RPprEpROHEkhv5K",
		"city_id": "153",
		"phone_number": "",
		"address": "Jl. Buntu",
		"image_url": ""
	}
}
```

##

### Delete User

Digunakan untuk menghapus pengguna berdasarkan ID.

#### Endpoint

```http
DELETE http://localhost:5000/api/user/adf4b794-f398-464a-b6d5-ef8a078f0705
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **adf4b794-f398-464a-b6d5-ef8a078f0705** (string,required): User ID.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success delete user with id 'adf4b794-f398-464a-b6d5-ef8a078f0705'"
}
```

##

### Find User

Digunakan untuk mencari pengguna berdasarkan kriteria tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/user/find?page=1&limit=5&search=a
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **page** (integer, optional): Nomor halaman yang diinginkan (opsional, default: 1).
- **limit** (integer, optional): Jumlah data per halaman (opsional, default: 10).
- **search** (string, optional): Kriteria pencarian untuk nama pengguna (opsional).

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"current_page": "1",
	"total_page": 2,
	"data": [
		{
			"id": "194e0179-febb-4bb3-a96b-b19001e0f21b",
			"username": "azhi",
			"email": "azhi@gmail.com",
			"phone_number": "",
			"address": "",
			"image_url": "",
			"created_at": "2023-10-13T20:49:08.295375+07:00",
			"updated_at": "2023-10-13T20:49:08.295375+07:00"
		},
		... // (data user lainnya)
	]
}
```

> [!NOTE]
> fitur ini hanya bisa digunakan untuk yang memiliki role admin selain admin maka akan mengembalikan pesan error unauthorized

##

### Get User

Digunakan untuk mendapatkan profil pengguna berdasarkan ID.

#### Endpoint

```http
GET http://localhost:5000/api/user/45313486-b690-4ab3-aa7d-86ef45be5628
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **45313486-b690-4ab3-aa7d-86ef45be5628** (string,required): User ID.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"id": "45313486-b690-4ab3-aa7d-86ef45be5628",
		"username": "azhi",
		"email": "azhi@gmail.com",
		"phone_number": "",
		"address": "",
		"image_url": "",
		"region": {
			"id": "153",
			"type": "Kota",
			"city_name": "Jakarta Selatan",
			"postal_code": "12230"
		},
		"store": {
			"id": "8bae8813-c361-4b1a-8c20-59f78010728e",
			"store_name": "Toko Buku",
			"description": "contoh deskripsi",
			"category": "pendidikan",
			"image_url": ""
		},
		"created_at": "2023-10-14T01:19:32.73704+07:00",
		"updated_at": "2023-10-16T12:16:57.062845+07:00"
	}
}
```

> [!NOTE]
> fitur ini hanya bisa digunakan untuk yang memiliki role admin selain admin maka akan mengembalikan pesan error unauthorized

##

### Create Store

Digunakan untuk membuat toko baru.

#### Endpoint

```http
POST http://localhost:5000/api/store/create
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Request Body

- **store_name** (string, required): Nama toko.
- **description** (string, required): Deskripsi toko.
- **category** (string, required): Kategori toko.
- **image_url** (string, optional): URL gambar toko (dapat dikosongkan).

### Contoh Request Body:

```json
{
	"store_name": "Toko Buku",
	"description": "contoh deskripsi",
	"category": "pendidikan",
	"image_url": ""
}
```

#### Response

- **HTTP Status**: 201 Created
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success Create Store"
}
```

##

### Update Store

Digunakan untuk memperbarui informasi toko yang ada.

#### Endpoint

```http
PUT http://localhost:5000/api/store/c15dc952-7fea-499c-b2cb-3c9d6fe8503a
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **c15dc952-7fea-499c-b2cb-3c9d6fe8503a** (string,required): Store ID.

### Request Body

- **store_name** (string, optional): Nama toko (dapat dikosongkan).
- **description** (string, optional): Deskripsi toko (dapat dikosongkan).
- **category** (string, optional): Kategori toko (dapat dikosongkan).
- **image_url** (string, optional): URL gambar toko (dapat dikosongkan).
- **address** (string,optional): address toko (dapat dikosongkan).
- **city_id** (string,optional): city id toko (dapat dikosongkan).

#### Contoh Request Body:

```json
{
	"store_name": "Toko Game",
	"description": "contoh deskripsi",
	"category": "hiburan",
	"image_url": "",
	"address": "",
	"city_id": "153"
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "Success Update Store With Id 'c15dc952-7fea-499c-b2cb-3c9d6fe8503a'",
	"data": {
		"store_name": "Toko Game",
		"description": "contoh deskripsi",
		"category": "hiburan",
		"image_url": "",
		"address": "",
		"city_id": "153"
	}
}
```

##

### Delete Store

Digunakan untuk menghapus toko berdasarkan ID.

#### Endpoint

```http
DELETE http://localhost:5000/api/store/cc57e8e1-ce13-45b8-ac87-a93ea8611294
```

#### Parameters

- **cc57e8e1-ce13-45b8-ac87-a93ea8611294** (string,required): Store ID.

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success Delete Store with id 'cc57e8e1-ce13-45b8-ac87-a93ea8611294'"
}
```

##

### Find Stores

Digunakan untuk mencari toko berdasarkan kriteria tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/store/find?page=1&limit=2&search=bu
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **page** (integer, optional): Nomor halaman yang diinginkan (opsional, default: 1).
- **limit** (integer, optional): Jumlah data per halaman (opsional, default: 10).
- **search** (string, optional): Kriteria pencarian untuk nama toko (opsional).

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
  "code": 200,
  "status": "OK",
  "current_page": "1",
  "total_page": 1,
  "data": [
    {
      "id": "8bae8813-c361-4b1a-8c20-59f78010728e",
      "store_name": "Toko Buku",
      "description": "contoh deskripsi",
      "category": "pendidikan",
      "image_url": "",
      "created_at": "2023-10-14T02:11:56.344211+07:00",
      "updated_at": "2023-10-14T02:38:15.126789+07:00"
    },
    ... // (data store lainnya)
  ]
}
```

##

### Get Store

Digunakan untuk mendapatkan profil toko berdasarkan ID.

#### Endpoint

```http
GET http://localhost:5000/api/store/8bae8813-c361-4b1a-8c20-59f78010728e
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **8bae8813-c361-4b1a-8c20-59f78010728e** (string,required): Store ID.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"id": "8bae8813-c361-4b1a-8c20-59f78010728e",
		"region": {
			"id": "153",
			"type": "Kota",
			"city_name": "Jakarta Selatan",
			"postal_code": "12230"
		},
		"store_name": "Toko Buku",
		"description": "contoh deskripsi",
		"category": "pendidikan",
		"image_url": "",
		"created_at": "2023-10-14T02:11:56.344211+07:00",
		"updated_at": "2023-10-14T02:38:15.126789+07:00"
	}
}
```

##

### Get Store Report

Mengambil laporan toko berdasarkan ID toko dengan rentang tanggal tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/store/report/8bae8813-c361-4b1a-8c20-59f78010728e?startDate=2023-10-22&endDate=2023-10-30
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **8bae8813-c361-4b1a-8c20-59f78010728e** (string,required): ID toko.
- **startDate**: Tanggal awal periode laporan.
- **endDate**: Tanggal akhir periode laporan.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"total_sales": 100000,
		"total_product_sold": 2
	}
}
```

##

### Create Product

Digunakan untuk membuat produk baru dalam toko.

#### Endpoint

```http
POST http://localhost:5000/api/product/create
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Request Body

- **product_name** (string, required): Nama produk.
- **description** (string, required): Deskripsi produk.
- **category** (string, required): Kategori produk.
- **stock** (integer, required): Jumlah stok produk.
- **price** (integer, required): Harga produk.
- **image_url** (string, optional): URL gambar produk (dapat dikosongkan).
- **weight_on_gram** (integer,required): berat product

#### Contoh Request Body:

```json
{
	"product_name": "Buku Politik",
	"description": "Contoh Deskripsi",
	"category": "Pendidikan",
	"stock": 99,
	"price": 50000,
	"image_url": "",
	"weight_on_gram": 1000
}
```

#### Response

- **HTTP Status**: 201 Created
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success Create Product"
}
```

##

### Update Product

Digunakan untuk memperbarui informasi produk yang ada.

#### Endpoint

```http
PUT http://localhost:5000/api/product/5aaa787d-02d3-4487-bc75-b543da26c897
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **5aaa787d-02d3-4487-bc75-b543da26c897** (string,required): Product ID.

#### Request Body

- **product_name** (string, optional): Nama produk (dapat dikosongkan).
- **description** (string, optional): Deskripsi produk (dapat dikosongkan).
- **category** (string, optional): Kategori produk (dapat dikosongkan).
- **stock** (integer, optional): Jumlah stok produk (dapat dikosongkan).
- **price** (integer, optional): Harga produk (dapat dikosongkan).
- **image_url** (string, optional): URL gambar produk (dapat dikosongkan).
- **weight_on_gram** (integer,optional): berat product (dapat dikosongkan).

#### Contoh Request Body:

```json
{
	"product_name": "Buku Politik",
	"description": "",
	"category": "",
	"stock": "",
	"price": "",
	"image_url": "",
	"weight_on_gram": 1000
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "Success Update Product With Id '5aaa787d-02d3-4487-bc75-b543da26c897'",
	"data": {
		"product_name": "Buku Politik",
		"description": "Contoh Deskripsi",
		"category": "Pendidikan",
		"stock": "99",
		"price": "50000",
		"image_url": "",
		"weight_on_gram": 1000
	}
}
```

##

### Delete Product

Digunakan untuk menghapus produk berdasarkan ID.

#### Endpoint

```http
DELETE http://localhost:5000/api/product/5aaa787d-02d3-4487-bc75-b543da26c897
```

#### Request Header

- **Access-Token** (string, optional): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **5aaa787d-02d3-4487-bc75-b543da26c897** (string,required): Product ID.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success Delete Product With Id '5aaa787d-02d3-4487-bc75-b543da26c897'"
}
```

##

### Find Products

Digunakan untuk mencari produk berdasarkan kriteria tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/product/find?search=a&page=1&limit=5
```

#### Request Header

- **Access-Token** (string,required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **page** (integer, optional): Nomor halaman yang diinginkan (opsional, default: 1).
- **limit** (integer, optional): Jumlah data per halaman (opsional, default: 10).
- **search** (string, optional): Kriteria pencarian untuk nama produk, category, & store ID (opsional).

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"current_page": "1",
	"total_page": 1,
	"data": [
		{
			"id": "5aaa787d-02d3-4487-bc75-b543da26c897",
			"product_name": "Buku Politik",
			"description": "Contoh Deskripsi",
			"category": "Pendidikan",
			"weight_on_gram": 1000,
			"stock": "99",
			"price": "50000",
			"image_url": "",
			"created_at": "2023-10-14T16:41:21.190663+07:00",
			"updated_at": "2023-10-14T19:26:54.898163+07:00"
		},
		... // (data product lainnya)
	]
}
```

##

### Get Product

Digunakan untuk mendapatkan rincian produk berdasarkan ID produk.

#### Endpoint

```http
GET http://localhost:5000/api/product/5aaa787d-02d3-4487-bc75-b543da26c897
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **5aaa787d-02d3-4487-bc75-b543da26c897** (string,required): Product ID.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"id": "5aaa787d-02d3-4487-bc75-b543da26c897",
		"store": {
			"id": "8bae8813-c361-4b1a-8c20-59f78010728e",
			"store_name": "Toko Buku",
			"category": "pendidikan",
			"image_url": ""
		},
		"product_name": "Buku Politik",
		"description": "Contoh Deskripsi",
		"category": "Pendidikan",
		"weight_on_gram": 1000,
		"stock": "99",
		"price": "50000",
		"image_url": "",
		"created_at": "2023-10-14T16:41:21.190663+07:00",
		"updated_at": "2023-10-14T19:26:54.898163+07:00"
	}
}
```

##

### Add Product to Cart

Digunakan untuk menambahkan produk ke keranjang belanja pengguna.

#### Endpoint

```http
POST http://localhost:5000/api/cart/add/5aaa787d-02d3-4487-bc75-b543da26c897?qty=2
```

### Request Header

- **Access-Token** (string,required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **5aaa787d-02d3-4487-bc75-b543da26c897** (string,required): ID produk yang ingin ditambahkan ke keranjang.
- **qty** (integer,optional): Jumlah produk yang ingin ditambahkan ke keranjang (opsional, default: 1).

#### Response

- **HTTP Status**: 201 Created
- **Content-Type**: application/json; charset=utf-8

```json
{
   "code": 201,
   "status": "Success Add Product With Id '5aaa787d-02d3-4487-bc75-b543da26c897' To Your Cart",
   "data": [
      {
         "cart_id": "177f95e3-2080-43d5-abdf-87030f313555",
         "store": {
            "store_name": "Toko Buku",
            "description": "contoh deskripsi",
            "category": "pendidikan",
            "image_url": ""
         },
         "items": [
            {
               "id":"8ff0034e-3507-43b0-8fee-6277c11347e0",
               "amount": "2",
               "product":{
                  "product_name": "Buku Politik",
                  "description": "Contoh Deskripsi",
                  "category": "Pendidikan",
                  "weight_on_gram":1000,
                  "price": "50000",
                  "image_url": ""
               }
            },... // (data item lainnya)
         ],
         "total_price": "100000",
         "total_gram": "1000",
         "created_at": "2023-10-17T10:22:59.801696+07:00",
         "updated_at": "2023-10-17T10:22:59.838224+07:00"
      },... // (data keranjang lainnya)
   ]
}
```

##

### Delete Cart

Digunakan untuk menghapus keranjang belanja pengguna.

#### Endpoint

```http
DELETE http://localhost:5000/api/cart/db1106eb-48be-49e4-9ebc-651331072944
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **db1106eb-48be-49e4-9ebc-651331072944** (string,required): ID keranjang yang ingin dihapus.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "Success Delete Cart With Id 'db1106eb-48be-49e4-9ebc-651331072944'",
	"data": []
}
```

##

### Update Item Cart

Digunakan untuk memperbarui jumlah produk dalam keranjang belanja.

#### Endpoint

```http
PUT http://localhost:5000/api/cart/item/f3aea3b7-e2be-47fd-b0a2-a3496537a3e1?qty=3
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **f3aea3b7-e2be-47fd-b0a2-a3496537a3e1** (string,required): ID item dalam keranjang yang ingin diperbarui.
- **qty** (integer,required): Jumlah produk yang ingin diperbarui dalam item ini.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
   "code": 200,
   "status": "Success Update Item 'f3aea3b7-e2be-47fd-b0a2-a3496537a3e1' With qty 3",
   "data": [
      {
         "cart_id": "109ed563-b36f-4784-8ff5-7b580da1939d",
         "store": {
            "store_name": "Toko Buku",
            "description": "contoh deskripsi",
            "category": "pendidikan",
            "image_url": ""
         },
         "items": [
            {
               "id": "f3aea3b7-e2be-47fd-b0a2-a3496537a3e1",
               "amount": "3",
               "product": {
                  "product_name": "Buku Politik",
                  "description": "Contoh Deskripsi",
                  "category": "Pendidikan",
                  "weight_on_gram":1000,
                  "price": "50000",
                  "image_url": ""
               }
            },... // (data item lainnya)
         ],
         "total_price": "150000",
         "total_gram": "3000",
         "created_at": "2023-10-18T21:07:28.338832+07:00",
         "updated_at": "2023-10-18T21:28:06.109172+07:00"
      },... // (data keranjang lainnya)
   ]
}
```

##

### Delete Item Cart

Digunakan untuk menghapus item tertentu dari keranjang belanja.

#### Endpoint

```http
DELETE http://localhost:5000/api/cart/item/8239ad03-ea2c-4ba5-b2a5-e3360c0d8ac0
```

#### Request Header

- **Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

#### Parameters

- **8239ad03-ea2c-4ba5-b2a5-e3360c0d8ac0** (string): ID item dalam keranjang yang ingin dihapus.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "Success Delete Item With Id '8239ad03-ea2c-4ba5-b2a5-e3360c0d8ac0'",
	"data": []
}
```

##

### Get Cart

Mengambil keranjang belanja pengguna saat ini.

#### Endpoint

```http
GET http://localhost:5000/api/cart/get
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
   "code": 200,
   "status": "OK",
   "data": [
      {
         "cart_id": "682f3ca8-9519-4b05-8463-700a383264cb",
         "store": {
            "store_name": "Toko Anime",
            "description": "contoh deskripsi",
            "category": "hiburan",
            "image_url": ""
         },
         "items":[
            {
               "id":"c162d728-3c49-4a8f-9d8a-a90dc86137b9",
               "amount": "2",
               "product":{
                  "product_name": "Komik Naruto",
          	  "description": "Contoh Deskripsi",
          	  "category": "Pendidikan",
                  "weight_on_gram":1000,
                  "price": "50000",
          	  "image_url": ""
               }
            },... // (data item lainnya)
         ],
         "total_price": "100000",
         "total_gram": "2000",
         "created_at": "2023-10-21T05:56:47.377428+07:00",
         "updated_at": "2023-10-21T05:56:47.38139+07:00"
      },... // (data keranjang lainnya)
   ]
}
```

##

### Cek Ongkir

Menghitung biaya pengiriman (ongkir) antara alamat pengiriman dan alamat tujuan menggunakan layanan ekspedisi tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/transaction/ongkir/e39eecaa-828e-45b1-9447-56aae81a8fe7?expedition=jne
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **e39eecaa-828e-45b1-9447-56aae81a8fe7** (string,required): ID keranjang belanja.
- **expedition** (string,required) : Nama ekspedisi yang digunakan untuk menghitung ongkir.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "OK",
	"data": {
		"origin_details": {
			"city_id": "151",
			"city_name": "Jakarta Barat",
			"postal_code": "11220",
			"province": "DKI Jakarta",
			"province_id": "6",
			"type": "Kota"
		},
		"destination_details": {
			"city_id": "153",
			"city_name": "Jakarta Selatan",
			"postal_code": "12230",
			"province": "DKI Jakarta",
			"province_id": "6",
			"type": "Kota"
		},
		"weight_on_gram": "2000",
		"services": [
			{
				"service": "CTC",
				"description": "JNE City Courier",
				"value": 20000,
				"etd": "1-2",
				"note": ""
			},
			{
				"service": "CTCYES",
				"description": "JNE City Courier",
				"value": 36000,
				"etd": "1-1",
				"note": ""
			},... // (data service lainnya)
		]
	}
}
```

> [!NOTE]
> untuk expedition hanya tersedia option jne, pos, tiki, menggunakan selain itu akan terkena error BAD REQUEST

##

### Checkout

Melakukan proses checkout untuk menyelesaikan pembelian dengan memilih metode pembayaran dan menghitung biaya pengiriman.

#### Endpoint

```http
POST http://localhost:5000/api/transaction/checkout/e39eecaa-828e-45b1-9447-56aae81a8fe7?payment=bca
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **e39eecaa-828e-45b1-9447-56aae81a8fe7** (string,required): ID keranjang belanja.
- **payment** (string,required): Metode pembayaran yang digunakan

#### Request Body

```json
{
	"origin_city": "Jakarta Barat",
	"destination_city": "Jakarta Selatan",
	"courier": "tiki",
	"weight_on_gram": "2000",
	"service": "ECO",
	"description": "Economy Service",
	"price": 16000
}
```

#### Response

- **HTTP Status**: 201 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"transaction_id": "8d2449b6-a26b-4b98-994f-df86da92134f",
	"order_id": "TRX-e39eecaa",
	"gross_amount": "116000.00",
	"payment_type": "bank_transfer",
	"transaction_time": "2023-10-23 00:29:18",
	"transaction_status": "pending",
	"fraud_status": "accept",
	"masked_card": "",
	"status_code": "201",
	"bank": "",
	"status_message": "Success, Bank Transfer transaction is created",
	"approval_code": "",
	"channel_response_code": "",
	"channel_response_message": "",
	"currency": "IDR",
	"card_type": "",
	"redirect_url": "",
	"id": "",
	"validation_messages": null,
	"installment_term": "",
	"eci": "",
	"saved_token_id": "",
	"saved_token_id_expired_at": "",
	"point_redeem_amount": 0,
	"point_redeem_quantity": 0,
	"point_balance_amount": "",
	"permata_va_number": "",
	"va_numbers": [
		{
			"bank": "bca",
			"va_number": "77855188791"
		}
	],
	"bill_key": "",
	"biller_code": "",
	"acquirer": "",
	"actions": null,
	"payment_code": "",
	"store": "",
	"qr_string": "",
	"on_us": false,
	"three_ds_version": "",
	"expiry_time": "2023-10-24 00:29:17"
}
```

> [!NOTE]
> Untuk metode pembayaran saat ini hanya tersedia bca, bni, dan bri selain itu maka akan mengembalikan response BAD REQUEST

##

### Get Transaction By User Id

Mengambil Semua transaksi berdasarkan ID pengguna.

#### Endpoint

```http
GET http://localhost:5000/api/transaction/user/4e9497c1-fb82-4301-94a8-7d0b4d6c4f53
```

#### Request Header

- **Access-Token**: Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **4e9497c1-fb82-4301-94a8-7d0b4d6c4f53** (string,required): User ID.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": "5654e222-c384-4dff-b8b2-f0cfe3a49914",
      "user": {
        "id": "4e9497c1-fb82-4301-94a8-7d0b4d6c4f53",
        "username": "sendi",
        "email": "sendi@gmail.com",
        "phone_number": "",
        "address": "",
        "image_url": "",
        "region": {
          "id": "153",
          "type": "Kota",
          "city_name": "Jakarta Selatan",
          "postal_code": "12230"
        }
      },
      "store": {
        "id": "8bae8813-c361-4b1a-8c20-59f78010728e",
        "region": {
          "id": "151",
          "type": "Kota",
          "city_name": "Jakarta Barat",
          "postal_code": "11220"
        },
        "store_name": "Toko Buku",
        "description": "contoh deskripsi",
        "category": "pendidikan",
        "image_url": "",
        "created_at": "2023-10-14T02:11:56.344211+07:00",
        "updated_at": "2023-10-14T02:38:15.126789+07:00"
      },
      "item": [
        {
          "id": "3f281685-16e9-40cf-a54c-9dea65853d5c",
          "product": {
            "product_name": "Komik Detektif Conan",
            "description": "Contoh Deskripsi",
            "category": "Pendidikan",
            "weight_on_gram": 1000,
            "price": "50000",
            "image_url": ""
          },
          "amount": "2"
        },... // (data item lainnya)
      ],
      "expedition": {
        "id": "ca36ca9d-c7d0-4efb-9c03-54ac235a37cf",
        "origin_city": "Jakarta Barat",
        "destination_city": "Jakarta Selatan",
        "courier": "tiki",
        "weight_on_gram": "2000",
        "service": "ECO",
        "description": "Economy Service",
        "price": 16000
      },
      "transaction_status": "pending",
      "total_product_price": "100000",
      "total_price": "116000",
      "created_at": "2023-10-23T00:29:20.255491+07:00",
      "updated_at": "2023-10-23T00:29:20.255491+07:00"
    },... // (data transaksi lainnya)
  ]
}
```

##

### Get Transaction By Store Id

Mengambil Semua transaksi berdasarkan ID Toko.

#### Endpoint

```http
GET http://localhost:5000/api/transaction/store/8bae8813-c361-4b1a-8c20-59f78010728e
```

#### Request Header

- **Access-Token**: Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **8bae8813-c361-4b1a-8c20-59f78010728e** (string,required): Store ID.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": "5654e222-c384-4dff-b8b2-f0cfe3a49914",
      "user": {
        "id": "4e9497c1-fb82-4301-94a8-7d0b4d6c4f53",
        "username": "sendi",
        "email": "sendi@gmail.com",
        "phone_number": "",
        "address": "",
        "image_url": "",
        "region": {
          "id": "153",
          "type": "Kota",
          "city_name": "Jakarta Selatan",
          "postal_code": "12230"
        }
      },
      "store": {
        "id": "8bae8813-c361-4b1a-8c20-59f78010728e",
        "region": {
          "id": "151",
          "type": "Kota",
          "city_name": "Jakarta Barat",
          "postal_code": "11220"
        },
        "store_name": "Toko Buku",
        "description": "contoh deskripsi",
        "category": "pendidikan",
        "image_url": "",
        "created_at": "2023-10-14T02:11:56.344211+07:00",
        "updated_at": "2023-10-14T02:38:15.126789+07:00"
      },
      "item": [
        {
          "id": "3f281685-16e9-40cf-a54c-9dea65853d5c",
          "product": {
            "product_name": "Komik Detektif Conan",
            "description": "Contoh Deskripsi",
            "category": "Pendidikan",
            "weight_on_gram": 1000,
            "price": "50000",
            "image_url": ""
          },
          "amount": "2"
        },... // (data item lainnya)
      ],
      "expedition": {
        "id": "ca36ca9d-c7d0-4efb-9c03-54ac235a37cf",
        "origin_city": "Jakarta Barat",
        "destination_city": "Jakarta Selatan",
        "courier": "tiki",
        "weight_on_gram": "2000",
        "service": "ECO",
        "description": "Economy Service",
        "price": 16000
      },
      "transaction_status": "pending",
      "total_product_price": "100000",
      "total_price": "116000",
      "created_at": "2023-10-23T00:29:20.255491+07:00",
      "updated_at": "2023-10-23T00:29:20.255491+07:00"
    },... // (data transaksi lainnya)
  ]
}
```

##

### Create Product Review

Membuat ulasan produk baru untuk produk tertentu.

#### Endpoint

```http
POST http://localhost:5000/api/review/create/190fce26-fd07-48df-9283-84759abbefae
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **190fce26-fd07-48df-9283-84759abbefae** (string,required): ID produk yang akan ditinjau.

#### Request Body

```json
{
	"comment": "Product bagus dan murah"
}
```

#### Response

- **HTTP Status**: 201 Created
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success Create Product Review"
}
```

##

### Get Product Review

Mengambil ulasan produk untuk produk tertentu.

#### Endpoint

```http
GET http://localhost:5000/api/review/get/190fce26-fd07-48df-9283-84759abbefae?page=1&limit=10
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **190fce26-fd07-48df-9283-84759abbefae** (string,required): ID produk yang ulasan produknya akan diambil.
- **page** (integer,optional): Nomor halaman yang diinginkan dalam hasil ulasan (opsional, default: 1).
- **limit** (integer,optional): Jumlah ulasan yang ingin ditampilkan dalam satu halaman (opsional, default: 10).

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
  "code": 200,
  "status": "OK",
  "current_page": "1",
  "total_page": 1,
  "data": [
    {
      "id": "c18b0b9e-fcf4-4863-8f33-f361f95c24f1",
      "comment": "Product bagus dan murah",
      "created_at": "2023-10-25T00:44:10.92995+07:00",
      "updated_at": "2023-10-25T00:44:10.92995+07:00"
    }... // (data product review lainnya)
  ]
}
```

##

### Update Product Review

Mengubah ulasan produk yang sudah ada.

#### Endpoint

```http
PUT http://localhost:5000/api/review/c18b0b9e-fcf4-4863-8f33-f361f95c24f1
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **c18b0b9e-fcf4-4863-8f33-f361f95c24f1** (string,required): ID ulasan yang akan di ubah.

#### Request Body

```json
{
	"comment": "Product bagus dan murah banget"
}
```

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"code": 200,
	"status": "Success Update Review With Id 'c18b0b9e-fcf4-4863-8f33-f361f95c24f1'",
	"data": {
		"comment": "Product bagus dan murah banget"
	}
}
```

##

### Delete Product Review

Menghapus ulasan produk.

#### Endpoint

```http
DELETE http://localhost:5000/api/review/c18b0b9e-fcf4-4863-8f33-f361f95c24f1
```

#### Request Header

- **Access-Token** (string,required): Token akses yang digunakan untuk mengidentifikasi pengguna yang terautentikasi.

#### Parameters

- **c18b0b9e-fcf4-4863-8f33-f361f95c24f1** (string,required): ID ulasan yang akan di hapus.

#### Response

- **HTTP Status**: 200 OK
- **Content-Type**: application/json; charset=utf-8

```json
{
	"msg": "Success Delete Review with id 'c18b0b9e-fcf4-4863-8f33-f361f95c24f1'"
}
```

## MIT License

Copyright (c) 2023 NURSANDY IHKSAN

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
