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
		"id": "1d7828b5-594a-4648-815d-ecd89b3e79b8",
		"username": "person",
		"email": "person@gmail.com",
		"password": "$2a$10$UbHxmNYQYPwDSk2rcKA1KeAY8vHywEpgSysxD61OvlB8JLszO4IwS",
		"phone_number": "",
		"address": "",
		"image_url": "",
		"created_at": "2023-10-12T18:26:01.950373+07:00",
		"updated_at": "2023-10-12T18:26:01.950373+07:00"
	}
}
```

> [!NOTE]
> Kata sandi (password) yang digunakan dalam contoh di atas telah di-hash dengan algoritma bcrypt.
> jika user memasukkan email dan no telephone yang sudah ada maka akan mengembalikan pesan error.

###
