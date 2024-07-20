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
      <td>`/register`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **Nama**: string [default ""]<br>
        - **Nim**: string [default "11"]<br>
        - **Password**: string [default "6"]<br>
        - **Cpassword**: string [default ""]
      </td>
      <td>
          <strong>Code:</strong> 200 OK<br>
          <strong>Data:</strong> [<br>
              {"id": "1",<br>
               "nama": "Muhammad Fathan Mukhlisan",<br>
               "nim": "09021182227009",<br>
               "password": "$2a$10$FSWf1TGdn0k0wi4IVxttxOoy<br>
                          31OSCWgDV8aryaZASa05dsDE5V8ou",<br>
                }]<br>
          <strong>View:</strong>
          ```register.html
          <p>Registrasi berhasil</p>
          ```
      </td>
    </tr>
    <tr>
      <td><strong>Login</strong></td>
      <td>`/login`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **nim**: string [default "11"]<br>
        - **password**: string [default "6"]
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>Headers:</strong> `Location: /dashboard`<br>
        <br>
        <strong>Success:</strong><br>
        <strong>Data:</strong> [<br>
              {"nim": "09021182227009",<br>
               "password": "testing123",<br>
         <strong>View:</strong> index.html (/dashboard)<br>
        <br>
        <strong>Failed:</strong><br>
        ```login.html<br>
        <p>NIM/Password Salah</p>
        ```
      </td>
    </tr>
    <tr>
      <td><strong>Profil</strong></td>
      <td>`/new-mahasiswa`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>View:</strong> profil.html
      </td>
    </tr>
    <tr>
      <td><strong>Profil</strong></td>
      <td>`/new-mahasiswa`</td>
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
        <strong>Code:</strong> 301 Moved Permanently<br>
        <strong>View:</strong> index.html<br>
        <strong>Data:</strong> [<br>
              {"no_reg": "1",<br>
               "nama": "Muhammad Fathan Mukhlisan",<br>
               "nim": "09021182227009",<br>
               "alamat": "Jln. White Rabbit",<br>
                "no": "081280818081"
                }]<br>
      </td>
    </tr>
    <tr>
      <td><strong>Dosen</strong></td>
      <td>`/new-dosen`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>View:</strong> dosen.html
      </td>
    </tr>
    <tr>
      <td><strong>Dosen</strong></td>
      <td>`/new-dosen`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **no_reg**: string [default ""]<br>
        - **nama_dosen**: string [default ""]<br>
        - **nip**: string [default ""]
      </td>
      <td>
        <strong>Code:</strong> 301 Moved Permanently<br>
        <strong>View:</strong> index.html (/dashboard)<br>
        <strong>Data:</strong> [<br>
              {"no_reg": "1",<br>
               "nama_dosen": "Axel Christensen",<br>
               "nip": "123456789",<br>
                }]<br>
      </td>
    </tr>
    <tr>
      <td><strong>Edit</strong></td>
      <td>`/edit-update?type={type}&id={id}`</td>
      <td>GET</td>
      <td>
        <strong>Request Params:</strong><br>
        - **type**: string [default "bioData"]<br>
        - **id**: string [default ""]
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>View:</strong> edit.html
      </td>
    </tr>
    <tr>
      <td><strong>Edit</strong></td>
      <td>`/edit-update?type={type}&id={id}`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **type**: string [default "bioData"]<br>
        - **id**: string [default ""]<br>
        - **nama**: string [default ""] (untuk `bioData`)<br>
        - **nim**: string [default ""] (untuk `bioData`)<br>
        - **alamat**: string [default ""] (untuk `bioData`)<br>
        - **no_telp**: string [default ""] (untuk `bioData`)<br>
        - **nama_dosen**: string [default ""] (untuk `dosen`)<br>
        - **nip**: string [default ""] (untuk `dosen`)
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>View:</strong> index.html (/dashboard)
      </td>
    </tr>
    <tr>
      <td><strong>Upload</strong></td>
      <td>`/upload`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>View:</strong> upload.html
      </td>
    </tr>
    <tr>
      <td><strong>Upload</strong></td>
      <td>`/upload`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **myFile**: file [default ""]<br>
        - **no_reg**: string [default ""]
      </td>
      <td>
        <strong>Code:</strong> 301 Moved Permanently<br>
        <strong>View:</strong> index.html(/dashboard)<br>
        <strong>Data:</strong> [<br>
              {"no_reg": "1",<br>
               "myFile": "Proposal Skripsi.pdf",<br>
                }]<br>
      </td>
    </tr>
    <tr>
      <td><strong>Status</strong></td>
      <td>`/status`</td>
      <td>GET</td>
      <td>
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>View:</strong> status.html
      </td>
    </tr>
    <tr>
      <td><strong>Status</strong></td>
      <td>`/status`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **id**: string [default ""]
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>Konten:</strong> <br>
        ```status.html <br>
        Nama NIM DosenPembimbing Judul Proposal                                     
        ```
      </td>
    </tr>
    <tr>
      <td><strong>Delete</strong></td>
      <td>`/delete`</td>
      <td>POST</td>
      <td>
        <strong>Request Params:</strong><br>
        - **id**: string [default ""]
      </td>
      <td>
        <strong>Code:</strong> 200 OK<br>
        <strong>Konten:</strong> <br>
        ```status.html <br>
        <br>                                    
        ```
      </td>
    </tr>
  </tbody>
</table>
