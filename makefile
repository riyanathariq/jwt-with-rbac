# Makefile to generate RSA public and private keys in single-line base64 encoded format

# Variables
PRIVATE_KEY_FILE = private.pem
PUBLIC_KEY_FILE = public.pem
BASE64_PRIVATE_KEY = base64_private_key.txt
BASE64_PUBLIC_KEY = base64_public_key.txt
KEY_SIZE = 2048

# start
start:
	@go run main.go

# Default target
all: generate-keys clean-intermediate

# Target to generate keys
generate-keys: $(BASE64_PRIVATE_KEY) $(BASE64_PUBLIC_KEY)

$(PRIVATE_KEY_FILE):
	@echo "Generating private key..."
	openssl genpkey -algorithm RSA -out $(PRIVATE_KEY_FILE) -pkeyopt rsa_keygen_bits:$(KEY_SIZE)

$(PUBLIC_KEY_FILE): $(PRIVATE_KEY_FILE)
	@echo "Generating public key..."
	openssl rsa -pubout -in $(PRIVATE_KEY_FILE) -out $(PUBLIC_KEY_FILE)

$(BASE64_PRIVATE_KEY): $(PRIVATE_KEY_FILE)
	@echo "Encoding private key in base64 (single line)..."
	openssl base64 -A -in $(PRIVATE_KEY_FILE) -out $(BASE64_PRIVATE_KEY)
	@echo "Private key (base64 encoded):"
	@cat $(BASE64_PRIVATE_KEY)

$(BASE64_PUBLIC_KEY): $(PUBLIC_KEY_FILE)
	@echo "Encoding public key in base64 (single line)..."
	openssl base64 -A -in $(PUBLIC_KEY_FILE) -out $(BASE64_PUBLIC_KEY)
	@echo "Public key (base64 encoded):"
	@cat $(BASE64_PUBLIC_KEY)

# Target to clean the intermediate .pem files
clean-intermediate: $(PRIVATE_KEY_FILE) $(PUBLIC_KEY_FILE)
	@echo "Cleaning up intermediate files..."
	rm -f $(PRIVATE_KEY_FILE) $(PUBLIC_KEY_FILE)

# Target to clean all generated files
clean:
	@echo "Cleaning up all generated files..."
	rm -f $(PRIVATE_KEY_FILE) $(PUBLIC_KEY_FILE) $(BASE64_PRIVATE_KEY) $(BASE64_PUBLIC_KEY)

# Phony targets
.PHONY: all generate-keys clean clean-intermediate
