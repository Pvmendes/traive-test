#!/bin/bash

# Set the base URL of your API
BASE_URL="http://localhost:8080"

# Function to make a POST request
function createTransaction() {
  local id="$1"
  local origin="$2"
  local userId="$3"
  local amount="$4"
  local operation="$5"
  local createdAt="$6"

  curl -X POST \
    -H "Content-Type: application/json" \
    -d '{        
      "id": "'"$id"'",
      "origin": "'"$origin"'",
      "user_id": "'"$userId"'",
      "amount": '"$amount"',
      "operation": "'"$operation"'",
      "created_at": "'"$createdAt"'"
    }' \
    "$BASE_URL/transactions/CreateTransaction"
}

# Number of transactions to create
NUM_TRANSACTIONS1=500

# Loop to create transactions
for ((i = 1; i <= NUM_TRANSACTIONS1; i++)); do
  id="$i"
  origin="desktop-web"
  userId=$((1000 + i))  # Auto-increment user_id
  amount=$((10 + i))
  operation="credit"  
  createdAt=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
  echo
  echo -e origin
  # Call the createTransaction function with the data
  createTransaction "$id" "$origin" "$userId" "$amount" "$operation" "$createdAt"

done

# Number of transactions to create
NUM_TRANSACTIONS2=1000

# Loop to create transactions
for ((i = NUM_TRANSACTIONS1; i <= NUM_TRANSACTIONS2; i++)); do
  id="$i"
  origin="mobile-android"
  userId=$((1000 + i))  # Auto-increment user_id
  amount=$((10 + i))
  operation="credit"  
  createdAt=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
  echo
  echo -e origin
  # Call the createTransaction function with the data
  createTransaction "$id" "$origin" "$userId" "$amount" "$operation" "$createdAt"
done

# Number of transactions to create
NUM_TRANSACTIONS3=1500

# Loop to create transactions
for ((i = NUM_TRANSACTIONS2; i <= NUM_TRANSACTIONS3; i++)); do
  id="$i"
  origin="mobile-ios"
  userId=$((1000 + i))  # Auto-increment user_id
  amount=$((10 + i))
  operation="credit"  
  createdAt=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
  echo
  echo -e origin
  # Call the createTransaction function with the data
  createTransaction "$id" "$origin" "$userId" "$amount" "$operation" "$createdAt"
done

