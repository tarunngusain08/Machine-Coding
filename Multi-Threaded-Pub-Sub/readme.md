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
