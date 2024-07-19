# Finito
## Endpoints
<table>
  <thead>
    <tr>
      <th>Feature</th>
      <th>Endpoint</th>
      <th>Method</th>
      <th>Request</th>
      <th>Response</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><strong>Register</strong></td>
      <td>`/register`</td>
      <td>GET</td>
      <td>
        <strong>Tujuan:</strong> Menampilkan formulir pendaftaran.<br>
        <strong>Response:</strong> Merender template `register.html`.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong><br>
        ```html<br>
        <p>Registrasi berhasil</p><br>
        ```
      </td>
    </tr>
    <tr>
      <td><strong>Register</strong></td>
      <td>`/register`</td>
      <td>POST</td>
      <td>
        <strong>Tujuan:</strong> Memproses pengiriman formulir pendaftaran.<br>
        <strong>Headers:</strong> `Content-Type: application/x-www-form-urlencoded`<br>
        <strong>Parameter Body:</strong><br>
        - `Nama` (string): Nama pengguna.<br>
        - `Nim` (string): Nomor ID pengguna.<br>
        - `Password` (string): Kata sandi pengguna.<br>
        - `Cpassword` (string): Konfirmasi kata sandi pengguna.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong><br>
        ```html<br>
        <p>Registrasi berhasil</p><br>
        ```
      </td>
    </tr>
    <tr>
      <td><strong>Login</strong></td>
      <td>`/login`</td>
      <td>GET</td>
      <td>
        <strong>Tujuan:</strong> Menampilkan formulir login.<br>
        <strong>Response:</strong> Merender template `login.html`.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir login.
      </td>
    </tr>
    <tr>
      <td><strong>Login</strong></td>
      <td>`/login`</td>
      <td>POST</td>
      <td>
        <strong>Tujuan:</strong> Memproses pengiriman formulir login.<br>
        <strong>Headers:</strong> `Content-Type: application/x-www-form-urlencoded`<br>
        <strong>Parameter Body:</strong><br>
        - `nim` (string): Nomor ID pengguna.<br>
        - `password` (string): Kata sandi pengguna.
      </td>
      <td>
        <strong>Kode:</strong> 303 See Other<br>
        <strong>Konten:</strong> Redirect `/dashboard`<br>
        <strong>Headers:</strong> `Location: /dashboard`<br>
        <strong>Respon Error:</strong><br>
        **Kode:** 200 OK<br>
        **Konten:**<br>
        ```html<br>
        <p>NIM/Password Salah</p><br>
        ```
      </td>
    </tr>
    <tr>
      <td><strong>Create(Profil)</strong></td>
      <td>`/new-mahasiswa`</td>
      <td>GET</td>
      <td>
        <strong>Tujuan:</strong> Menampilkan formulir pembuatan profil baru.<br>
        <strong>Response:</strong> Merender template `profil.html`.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir pembuatan profil.
      </td>
    </tr>
    <tr>
      <td><strong>Create(Profil)</strong></td>
      <td>`/new-mahasiswa`</td>
      <td>POST</td>
      <td>
        <strong>Tujuan:</strong> Memproses pengiriman formulir untuk membuat data mahasiswa baru di database.<br>
        <strong>Headers:</strong> `Content-Type: application/x-www-form-urlencoded`<br>
        <strong>Parameter Body:</strong><br>
        - `no_reg` (string): Nomor registrasi mahasiswa.<br>
        - `nama` (string): Nama mahasiswa.<br>
        - `nim` (string): Nomor ID mahasiswa.<br>
        - `alamat` (string): Alamat mahasiswa.<br>
        - `no` (string): Nomor telepon mahasiswa.
      </td>
      <td>
        <strong>Kode:</strong> 301 Moved Permanently<br>
        <strong>Konten:</strong> Redirect `/dashboard`<br>
        <strong>Respon Error Internal Server:</strong><br>
        **Kode:** 500 Internal Server Error<br>
        **Konten:**<br>
        ```json<br>
        {<br>
        "error": "Pesan kesalahan server internal"<br>
        }<br>
        ```
      </td>
    </tr>
    <tr>
      <td><strong>Create(Dosen)</strong></td>
      <td>`/new-dosen`</td>
      <td>GET</td>
      <td>
        <strong>Tujuan:</strong> Menampilkan formulir untuk membuat dosen baru.<br>
        <strong>Response:</strong> Merender template `dosen.html`.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir.
      </td>
    </tr>
    <tr>
      <td><strong>Create(Dosen)</strong></td>
      <td>`/new-dosen`</td>
      <td>POST</td>
      <td>
        <strong>Tujuan:</strong> Memproses pengiriman formulir untuk membuat data dosen baru.<br>
        <strong>Headers:</strong> `Content-Type: application/x-www-form-urlencoded`<br>
        <strong>Parameter Body:</strong><br>
        - `no_reg` (string): Nomor registrasi dosen.<br>
        - `nama_dosen` (string): Nama dosen.<br>
        - `nip` (string): Nomor ID dosen.
      </td>
      <td>
        <strong>Kode:</strong> 301 Moved Permanently<br>
        <strong>Konten:</strong> Redirect `/dashboard`.
      </td>
    </tr>
    <tr>
      <td><strong>Edit</strong></td>
      <td>`/edit-update?type={type}&id={id}`</td>
      <td>GET</td>
      <td>
        <strong>Tujuan:</strong> Menampilkan formulir untuk mengedit data berdasarkan jenis entitas.<br>
        <strong>Headers:</strong> `Content-Type: application/x-www-form-urlencoded`<br>
        <strong>Parameter Query:</strong><br>
        - `type` (string): Jenis entitas (`biodata` atau `dosen`).<br>
        - `id` (string): ID data yang akan diedit.<br>
        <strong>Response:</strong> Merender template `edit.html` dengan data entitas.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender formulir dengan data yang ada.
      </td>
    </tr>
    <tr>
      <td><strong>Edit</strong></td>
      <td>`/edit-update?type={type}&id={id}`</td>
      <td>POST</td>
      <td>
        <strong>Tujuan:</strong> Memproses pengiriman formulir untuk memperbarui data berdasarkan jenis entitas.<br>
        <strong>Headers:</strong> `Content-Type: application/x-www-form-urlencoded`<br>
        <strong>Parameter Query:</strong><br>
        - `type` (string): Jenis entitas (`biodata` atau `dosen`).<br>
        - `id` (string): ID data yang akan diperbarui.<br>
        <strong>Parameter Body:</strong><br>
        - Untuk tipe `biodata`:<br>
          - `nama` (string): Nama individu.<br>
          - `nim` (string): Nomor ID.<br>
          - `alamat` (string): Alamat.<br>
          - `no_telp` (string): Nomor telepon.<br>
        - Untuk tipe `dosen`:<br>
          - `nama_dosen` (string): Nama dosen.<br>
          - `nip` (string): Nomor ID dosen.
      </td>
      <td>
        <strong>Kode:</strong> 303 See Other<br>
        <strong>Konten:</strong> Redirect `/dashboard`.
      </td>
    </tr>
    <tr>
      <td><strong>Upload</strong></td>
      <td>`/upload`</td>
      <td>GET</td>
      <td>
        <strong>Tujuan:</strong> Menampilkan formulir unggah file.<br>
        <strong>Response:</strong> Merender template `proposal.html`.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir unggah file.
      </td>
    </tr>
    <tr>
      <td><strong>Upload</strong></td>
      <td>`/upload`</td>
      <td>POST</td>
      <td>
        <strong>Tujuan:</strong> Memproses unggah file dan menyimpan file di database.<br>
        <strong>Headers:</strong> `Content-Type: multipart/form-data`<br>
        <strong>Parameter Body:</strong><br>
        - `myFile` (file): File yang akan diunggah.<br>
        - `no_reg` (string): Nomor registrasi yang terkait dengan file.
      </td>
      <td>
        <strong>Kode:</strong> 301 Moved Permanently<br>
        <strong>Konten:</strong> Redirect `/dashboard`.
      </td>
    </tr>
    <tr>
      <td><strong>Status</strong></td>
      <td>`/status`</td>
      <td>GET</td>
      <td>
        <strong>Tujuan:</strong> Menampilkan ringkasan data dari berbagai tabel.<br>
        <strong>Response:</strong> Merender template `status.html` dengan data ringkasan.
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML data ringkasan.
      </td>
    </tr>
    <tr>
      <td><strong>Delete</strong></td>
      <td>`/delete`</td>
      <td>POST</td>
      <td>
        <strong>Tujuan:</strong> Menghapus data berdasarkan ID yang diberikan.<br>
        <strong>Headers:</strong> `Content-Type: application/x-www-form-urlencoded`<br>
        <strong>Parameter Body:</strong><br>
        - `id` (string): ID data yang akan dihapus.
      </td>
      <td>
        <strong>Kode:</strong> 303 See Other<br>
        <strong>Konten:</strong> Redirect `/status`.
      </td>
    </tr>
  </tbody>
</table>
