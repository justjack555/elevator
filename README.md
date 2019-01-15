# Elevator

### Summary

Elevator is an attempt at creating a smart elevator designed as a distributed system. The elevator consists of 1+ "masters", 1+ worker "elevators", and a calculation service. Masters respond to events, such as push of a button or a camera sensing a new elevator rider. The master then invokes the calculation service to determine the worker elevator to respond to the new rider. This project is in its early stages.

### Build and Run the system
1. Open a terminal and navigate to your preferred directory
2. `git clone https://github.com/justjack555/elevator.git`
3. To launch each of the services at once, run: `./runElevator`
    
### Tasks to do
- [x] Initialization of Master
- [x] Signatures for core RPC functions
- [ ] Add Elevator setup at server initialization
- [ ] Add parameter parsing in Elevator request calls to master
