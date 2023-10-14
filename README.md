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
- **phone_number** (string, optional): No telephone pengguna.
- **address** (string, optional): Alamat pengguna (dapat dikosongkan).
- **image_url** (string, optional): URL gambar profil pengguna (dapat dikosongkan).

#### Contoh Request Body:

```json
{
	"username": "person",
	"email": "person@gmail.com",
	"password": "123",
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
	"code": 201,
	"status": "Successful user registration",
	"data": {
		"username": "person",
		"email": "person@gmail.com",
		"password": "$2a$10$UbHxmNYQYPwDSk2rcKA1KeAY8vHywEpgSysxD61OvlB8JLszO4IwS",
		"phone_number": "",
		"address": "",
		"image_url": ""
	}
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

#### Request Body

- **username** (string, optional): Nama pengguna (dapat dikosongkan).
- **email** (string, optional): Alamat email pengguna (dapat dikosongkan).
- **password** (string, optional): Kata sandi pengguna (dapat dikosongkan).
- **phone_number** (string, optional): Nomor telepon pengguna (dapat dikosongkan).
- **address** (string, optional): Alamat pengguna (dapat dikosongkan).
- **image_url** (string, optional): URL gambar profil pengguna (dapat dikosongkan).

#### Contoh Request Body:

```json
{
	"username": "",
	"email": "person@gmail.com",
	"password": "",
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

**Access-Token** (string, required): Token akses yang sah untuk mengotentikasi pengguna.

**Query Parameters**

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
		...
	]
}
```

> [!NOTE]
> fitur ini hanya bisa digunakan untuk yang memiliki role admin selain admin maka akan mengembalikan pesan error

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
	"code": 201,
	"status": "Success Create Store",
	"data": {
		"store_name": "Toko Buku",
		"description": "contoh deskripsi",
		"category": "pendidikan",
		"image_url": ""
	}
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

### Request Body

- **store_name** (string, optional): Nama toko (dapat dikosongkan).
- **description** (string, optional): Deskripsi toko (dapat dikosongkan).
- **category** (string, optional): Kategori toko (dapat dikosongkan).
- **image_url** (string, optional): URL gambar toko (dapat dikosongkan).

#### Contoh Request Body:

```json
{
	"store_name": "Toko Game",
	"description": "contoh deskripsi",
	"category": "hiburan",
	"image_url": ""
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
		"image_url": ""
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
#### Query Parameters
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
    ...
  ]
}
```