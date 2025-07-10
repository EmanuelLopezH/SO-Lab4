#include <stdio.h>
#include <pthread.h>
#include <semaphore.h>
#include <unistd.h>

#define N 5
sem_t forks[N];
sem_t room;

void* philosopher(void* num) {
    int id = *(int*)num;
    while (1) {
        printf("Philosopher %d thinking...\n", id);
        sleep(1);

        sem_wait(&room);
        sem_wait(&forks[id]);
        sem_wait(&forks[(id+1)%N]);

        printf("Philosopher %d eating...\n", id);
        sleep(1);

        sem_post(&forks[id]);
        sem_post(&forks[(id+1)%N]);
        sem_post(&room);
    }
}

int main() {
    pthread_t phil[N];
    int ids[N];
    sem_init(&room, 0, N - 1);
    for (int i = 0; i < N; i++) {
        sem_init(&forks[i], 0, 1);
        ids[i] = i;
        pthread_create(&phil[i], NULL, philosopher, &ids[i]);
    }
    for (int i = 0; i < N; i++)
        pthread_join(phil[i], NULL);
    return 0;
}
