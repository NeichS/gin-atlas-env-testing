# gin-atlas-env-testing

## ent

Para crear nuevas entidades podes utilizar el siguiente comando

```Shell
# debemos estar en el directorio raiz del proyecto app/

go run -mod=mod entgo.io/ent/cmd/ent new Entidad1 Entidad2 EntidadN
```

Ahora regeneramos el esquema a partir de nuetras nuevas entidades

```Shell
go generate ./ent
```

> Van a ver que hay mucho boilerplate pero el UNICO directorio que nos interesa dentro de ent es schema, ahi es donde vamos a definir nuestras entidades.

Una vez realizado esto, podemos proceder a crear la migración.

Cabe destacar que no necesitamos regenerar el esquema cuando no creamos entidades nuevas, si hacemos una añadido/eliminacion/actualizacion de un modelo ya existente solamente migramos.

## Atlas

Creamos base de datos de desarrollo

```Shell
just create-dev-db
```

Preguntamos por el estado de la base de datos de la cual no queremos perder los datos.

```Shell
just status
```

Aplicamos migraciones a nuestra base de datos.

```Shell
just migrate
```

Si volvemos a realizar cambios en el esquema debemos limpiar la base de datos de desarrollo vieja para generar la diferencia de esquema la cual luego impacta en la base de datos principal.

```Shell
just drop-dev-db

just create-dev-db
```

Ahora creamos la migracion la cual genera el sql para poder actualizar nuestro esquema.

```Shell
just create <migration_name>
```

Con este comando podemos ver un diagnóstico de la migracion que acabamos de realizar.

```Shell
just lint
```

