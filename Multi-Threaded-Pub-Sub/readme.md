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


## Multi Goroutine Simulation

<img width="706" alt="Screenshot 2024-06-13 at 6 13 14 PM" src="https://github.com/tarunngusain08/Machine-Coding/assets/36428256/5c87e884-6c55-4034-8e3d-67eb635c5142">


## Validations and Screenshots

<img width="1396" alt="Screenshot 2024-06-13 at 6 10 39 PM" src="https://github.com/tarunngusain08/Machine-Coding/assets/36428256/732a6603-17d6-43f6-88d6-a5c5b08719a8">


<img width="1390" alt="Screenshot 2024-06-13 at 6 07 15 PM" src="https://github.com/tarunngusain08/Machine-Coding/assets/36428256/75ca98ef-086e-4dd2-b145-30ec1171551e">


<img width="1445" alt="Screenshot 2024-06-13 at 6 08 01 PM" src="https://github.com/tarunngusain08/Machine-Coding/assets/36428256/2aa009b8-b0ff-466c-bd40-c8b5ef08dcff">
