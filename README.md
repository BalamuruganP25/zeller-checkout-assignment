# Zeller Checkout Assignment

This project is a sample checkout system for Zeller's computer store, implemented in Go.  
It demonstrates flexible pricing rules and a simple, extensible architecture.

## Features

- **Product Catalogue:**  
  - Super iPad (`ipd`): $549.99  
  - MacBook Pro (`mbp`): $1399.99  
  - Apple TV (`atv`): $109.50  
  - VGA adapter (`vga`): $30.00

- **Pricing Rules:**  
  - **3 for 2 Deal:** Buy 3 Apple TVs, pay for 2  
  - **Bulk Discount:** Buy more than 4 Super iPads, price drops to $499.99 each  
  - **Default Pricing:** All other products use their standard price

- **Flexible Rule Engine:**  
  - Easily add or modify pricing rules

## How to Run This Project

1. **Run the main program:**

   ```sh
   make run
   ```

2. **Run the tests:**

   ```sh
   make test
   ```

3. **Build the binary:**

   ```sh
   make build
   ```

4. **Clean up build artifacts:**

   ```sh
   make clean
   ```