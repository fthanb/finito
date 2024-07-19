# FINITO

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
      <td>`http://localhost:8000/register`</td>
      <td>GET</td>
      <td>
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
      <td>`http://localhost:8000/register`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **Nama**: string [default ""]<br>
        - **Nim**: string [default ""]<br>
        - **Password**: string [default ""]<br>
        - **Cpassword**: string [default ""]
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
      <td>`http://localhost:8000/login`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir login.
      </td>
    </tr>
    <tr>
      <td><strong>Login</strong></td>
      <td>`http://localhost:8000/login`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **nim**: string [default ""]<br>
        - **password**: string [default ""]
      </td>
      <td>
        <strong>Kode:</strong> 303 See Other<br>
        <strong>Konten:</strong> Mengalihkan ke `/dashboard`<br>
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
      <td><strong>NewMahasiswa</strong></td>
      <td>`http://localhost:8000/new-mahasiswa`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir pembuatan profil.
      </td>
    </tr>
    <tr>
      <td><strong>NewMahasiswa</strong></td>
      <td>`http://localhost:8000/new-mahasiswa`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **no_reg**: string [default ""]<br>
        - **nama**: string [default ""]<br>
        - **nim**: string [default ""]<br>
        - **alamat**: string [default ""]<br>
        - **no**: string [default ""]
      </td>
      <td>
        <strong>Kode:</strong> 301 Moved Permanently<br>
        <strong>Konten:</strong> Mengalihkan ke `/dashboard`<br>
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
      <td><strong>NewDosen</strong></td>
      <td>`http://localhost:8000/new-dosen`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir.
      </td>
    </tr>
    <tr>
      <td><strong>NewDosen</strong></td>
      <td>`http://localhost:8000/new-dosen`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **no_reg**: string [default ""]<br>
        - **nama_dosen**: string [default ""]<br>
        - **nip**: string [default ""]
      </td>
      <td>
        <strong>Kode:</strong> 301 Moved Permanently<br>
        <strong>Konten:</strong> Mengalihkan ke `/dashboard`.
      </td>
    </tr>
    <tr>
      <td><strong>EditAndUpdate</strong></td>
      <td>`http://localhost:8000/edit-update?type={type}&id={id}`</td>
      <td>GET</td>
      <td>
        <strong>Request Params:</strong><br>
        - **type**: string [default "biodata"]<br>
        - **id**: string [default ""]
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender formulir dengan data yang ada.
      </td>
    </tr>
    <tr>
      <td><strong>EditAndUpdate</strong></td>
      <td>`http://localhost:8000/edit-update?type={type}&id={id}`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **type**: string [default "biodata"]<br>
        - **id**: string [default ""]<br>
        - **nama**: string [default ""] (untuk `biodata`)<br>
        - **nim**: string [default ""] (untuk `biodata`)<br>
        - **alamat**: string [default ""] (untuk `biodata`)<br>
        - **no_telp**: string [default ""] (untuk `biodata`)<br>
        - **nama_dosen**: string [default ""] (untuk `dosen`)<br>
        - **nip**: string [default ""] (untuk `dosen`)
      </td>
      <td>
        <strong>Kode:</strong> 303 See Other<br>
        <strong>Konten:</strong> Mengalihkan ke `/dashboard`.
      </td>
    </tr>
    <tr>
      <td><strong>Upload</strong></td>
      <td>`http://localhost:8000/upload`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML formulir unggah file.
      </td>
    </tr>
    <tr>
      <td><strong>Upload</strong></td>
      <td>`http://localhost:8000/upload`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **myFile**: file [default ""]<br>
        - **no_reg**: string [default ""]
      </td>
      <td>
        <strong>Kode:</strong> 301 Moved Permanently<br>
        <strong>Konten:</strong> Mengalihkan ke `/dashboard`.
      </td>
    </tr>
    <tr>
      <td><strong>Status</strong></td>
      <td>`http://localhost:8000/status`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Kode:</strong> 200 OK<br>
        <strong>Konten:</strong> Merender HTML data ringkasan.
      </td>
    </tr>
    <tr>
      <td><strong>Status</strong></td>
      <td>`http://localhost:8000/status`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **id**: string [default ""]
      </td>
      <td>
        <strong>Kode:</strong> 303 See Other<br>
        <strong>Konten:</strong> Mengalihkan ke `/status`.
      </td>
    </tr>
    <tr>
      <td><strong>DeleteHandler</strong></td>
      <td>`http://localhost:8000/delete`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **id**: string [default ""]
      </td>
      <td>
        <strong>Kode:</strong> 303 See Other<br>
        <strong>Konten:</strong> Mengalihkan ke `/status`.
      </td>
    </tr>
  </tbody>
</table>
