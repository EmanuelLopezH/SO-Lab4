#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>

#define SIZE 10

typedef struct {
    int items[SIZE];
    int front, rear, count;
    pthread_mutex_t lock;
    pthread_cond_t not_empty;
    pthread_cond_t not_full;
} ThreadSafeQueue;

void init_queue(ThreadSafeQueue *q) {
    q->front = q->rear = q->count = 0;
    pthread_mutex_init(&q->lock, NULL);
    pthread_cond_init(&q->not_empty, NULL);
    pthread_cond_init(&q->not_full, NULL);
}

void enqueue(ThreadSafeQueue *q, int item) {
    pthread_mutex_lock(&q->lock);
    while (q->count == SIZE)
        pthread_cond_wait(&q->not_full, &q->lock);

    q->items[q->rear] = item;
    q->rear = (q->rear + 1) % SIZE;
    q->count++;

    pthread_cond_signal(&q->not_empty);
    pthread_mutex_unlock(&q->lock);
}

int dequeue(ThreadSafeQueue *q) {
    pthread_mutex_lock(&q->lock);
    while (q->count == 0)
        pthread_cond_wait(&q->not_empty, &q->lock);

    int item = q->items[q->front];
    q->front = (q->front + 1) % SIZE;
    q->count--;

    pthread_cond_signal(&q->not_full);
    pthread_mutex_unlock(&q->lock);
    return item;
}

ThreadSafeQueue q;

void* producer(void* arg) {
    for (int i = 0; i < 20; i++) {
        printf("Producer enqueuing %d\n", i);
        enqueue(&q, i);
        usleep(100000);
    }
    return NULL;
}

void* consumer(void* arg) {
    for (int i = 0; i < 20; i++) {
        int val = dequeue(&q);
        printf("Consumer dequeued %d\n", val);
        usleep(150000);
    }
    return NULL;
}

int main() {
    init_queue(&q);
    pthread_t prod, cons;
    pthread_create(&prod, NULL, producer, NULL);
    pthread_create(&cons, NULL, consumer, NULL);
    pthread_join(prod, NULL);
    pthread_join(cons, NULL);
    return 0;
}
