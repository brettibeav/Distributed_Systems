# Laboratorio 2 - Grupo 13

## Integrantes
- Beatrice Valdes - 201941556-5
- Tomas Garreton - 201823565-2

# Ejecución
- `make docker`: compilar y ejecutar el servidor
- `make capitanes`: ejecutar el cliente
- `make clean`: limpiar los contenedores

# Consideraciones
- Asumimos que el botin encontrado se resta de los botines de los planetas.

# Informe
- Agregar una central que ejecute las mismas funciones que la actualmente implementada podría tener ventajas y desventajas. Dentro de las primeras habría mayor disponibilidad en caso de que uno de los dos servidores centrales falle y así no habrían interrupciones. De la misma manera se podrían distribuir la carga y existiría la posibilidad de escalar y agregar más servidores según sea necesario. Las desventajas van relacionadas al desafío que supone mantener, administrar y configurar dos servidores sincronizados, incluyendo los costos que esto supone en hardware y recursos en general.
