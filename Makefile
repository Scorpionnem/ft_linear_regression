BINS = train predict

all: $(BINS)

train:
	go build -o train ./cmd/train

predict:
	go build -o predict ./cmd/predict

clean:
	rm -f $(BINS) *.png

re: clean all

.PHONY: all clean re train predict
