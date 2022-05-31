# Servidor
###FAKE-JWT-TOKEN

POST: `/v1/login:` JSON: `{"username":"admin","password":"12345"}`

200:

```
{
    "id_agencia": 1,
    "usuario": {
        "password": null,
        "username": "admin",
        "authorities": [
            {
                "authority": "ROLE_ADMIN"
            }
        ],
        "accountNonExpired": true,
        "accountNonLocked": true,
        "credentialsNonExpired": true,
        "enabled": true
    },
    "recursos": [
        {
            "title": "Configuración",
            "icon": "fa fa-cog",
            "deploymentOrder": 1,
            "resourceLoginQueryDtoList": [
                {
                    "title": "Dominios",
                    "route": "/domains"
                }
            ]
        },
        {
            "title": "Usuarios",
            "icon": "fa fa-cog",
            "deploymentOrder": 2,
            "resourceLoginQueryDtoList": [
                {
                    "title": "Usuarios",
                    "route": "/users"
                },
                {
                    "title": "Roles",
                    "route": "/roles"
                },
                {
                    "title": "Menús",
                    "route": "/resources"
                }
            ]
        },
        {
            "title": "Reporte",
            "icon": "fa fa-cog",
            "deploymentOrder": 4,
            "resourceLoginQueryDtoList": [
                {
                    "title": "Reportes",
                    "route": "/reports"
                }
            ]
        }
    ],
    "mensaje": "Hola admin, has iniciado sesión con éxito!",
    "telefono": "292292",
    "permisos": [
        {
            "title": "Seleccione el acceso a la información permitida",
            "permissionGroupCode": "PGR01",
            "rolePermissionQueryDtoList": [
                {
                    "id": 1,
                    "title": "Reporte pdf",
                    "permissionCode": "PR001",
                    "assign": true
                },
                {
                    "id": 2,
                    "title": "Reporte csv",
                    "permissionCode": "PR002",
                    "assign": true
                }
            ]
        }
    ],
    "nombre": "ISRAEL GUILLERMO QUISPE MAMANI",
    "email": "iquispe@datec.com.bo",
    "agencia": "Oficina Central Banco Union ",
    "token": "eyJhbGciOiJIUzUxMiJ9.eyJhdXRob3JpdGllcyI6Ilt7XCJhdXRob3JpdHlcIjpcIlJPTEVfQURNSU5cIn1dIiwic3ViIjoiYWRtaW4iLCJpYXQiOjE2NTQwMjk2MzUsImV4cCI6MTY1NDA0MzYzNX0.LSgsI0ucjQ-u80KslU7Lmd70ugI1ShTungdhSSKqLtIXWOJb6WrsYpn2RD2bKxF6w2a0Zt-81Sy1YZ5IMclLLw"
}
```
