# Multi-Threaded Publish/Subscribe Messaging System

## Description

This project is a multi-threaded publish/subscribe (Pub/Sub) messaging system implemented in Go. It allows publishers to broadcast messages to multiple subscribers in a thread-safe manner using Go's goroutines and channels. The system is designed to handle concurrent operations efficiently, ensuring that messages are delivered correctly and promptly to all subscribers.

## Key Components

1. **Publisher and Subscriber Management**:
   - **Publisher**: A component that sends messages to a specific topic.
   - **Subscriber**: A component that receives messages from a subscribed topic.
   - The system maintains a registry of topics and the associated subscribers.

2. **Message Broadcasting to Subscribers**:
   - When a publisher sends a message to a topic, the message is broadcast to all subscribers of that topic.
   - Subscribers receive messages asynchronously, ensuring non-blocking communication.

3. **Thread-Safe Operation**:
   - Utilizes Go's goroutines and channels to manage concurrent operations.
   - Ensures that the system is thread-safe, preventing race conditions and ensuring data integrity.

## Validations and Screenshots

<img width="930" alt="Screenshot 2024-06-12 at 6 04 32 PM" src="https://github.com/tarunngusain08/Machine-Coding/assets/36428256/3014695a-fc8c-4a48-9f81-531269777be3">

<img width="1033" alt="Screenshot 2024-06-12 at 6 04 44 PM" src="https://github.com/tarunngusain08/Machine-Coding/assets/36428256/726b2f37-0b9c-4941-9db5-d0e8c2b8ad81">

<img width="893" alt="Screenshot 2024-06-12 at 6 04 55 PM" src="https://github.com/tarunngusain08/Machine-Coding/assets/36428256/533a5c27-a30f-4ebb-b9ea-a9d0886a4a93">
