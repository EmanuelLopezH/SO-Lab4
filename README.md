# Laboratorio 4 - Sistemas Operativos

### Universidad de Antioquia

## Autores

Emanuel López Higuita - Andrés Calvo Areiza

Este repositorio contiene las soluciones en **C** y **Go** a los tres ejercicios del laboratorio sobre mecanismos de sincronización:

---

## Ejercicios Implementados

1. **Cola segura (Thread-Safe Queue)**

   - Archivos: `queue.c`, `queue.go`
   - Usa `pthread_mutex` y `pthread_cond` en C; `sync.Mutex` y `sync.Cond` en Go.

2. **Productor-Consumidor con Semáforos**

   - Archivos: `producer_consumer.c`, `producer_consumer.go`
   - Utiliza semáforos y mutex para coordinar acceso al buffer.

3. **Filósofos comensales**
   - Archivos: `dining_philosophers.c`, `dining_philosophers.go`
   - Solución sin interbloqueo usando semáforo "room" y forks individuales.

---

## Compilación y Ejecución

### C

Asegurarse de tener instalado `gcc` y las librerías `pthread` y `semaphore`.

```bash
# Compilar
gcc -o queue queue.c -lpthread
gcc -o producer_consumer producer_consumer.c -lpthread
gcc -o dining_philosophers dining_philosophers.c -lpthread

# Ejecutar
./queue
./producer_consumer
./dining_philosophers
```

### Go

Asegurarse de tener instalado `Go`.

```bash
# Ejecutar
go run queue.go
go run producer_consumer.go
go run dining_philosophers.go
```

---
# SO-Lab4
