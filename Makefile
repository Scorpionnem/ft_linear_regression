CXX = c++
CXXFLAGS = -Wall -Wextra -Werror -MMD -MP -g -std=c++17

INCLUDES = -I includes

PREDICT_NAME = Predict
TRAIN_NAME = Train

PREDICT_SRCS =	src/predict/main.cpp src/estimatePrice.cpp
TRAIN_SRCS =	src/train/main.cpp src/estimatePrice.cpp

OBJDIR = obj
PREDICT_OBJS = $(PREDICT_SRCS:%.cpp=$(OBJDIR)/%.o)
PREDICT_DEPS = $(PREDICT_SRCS:%.cpp=$(OBJDIR)/%.d)

TRAIN_OBJS = $(TRAIN_SRCS:%.cpp=$(OBJDIR)/%.o)
TRAIN_DEPS = $(TRAIN_SRCS:%.cpp=$(OBJDIR)/%.d)

all: $(PREDICT_NAME) $(TRAIN_NAME)

matplot:
	git clone https://github.com/alandefreitas/matplotplusplus/ external/matplotplusplus
	cd external/matplotplusplus ; cmake -B build ; cmake --build build

re: fclean all

$(PREDICT_NAME): $(PREDICT_OBJS)
	@echo Compiling $(PREDICT_NAME)
	@$(CXX) $(CXXFLAGS) $(INCLUDES) -o $@ $^

$(TRAIN_NAME): $(TRAIN_OBJS)
	@echo Compiling $(TRAIN_NAME)
	@$(CXX) $(CXXFLAGS) $(INCLUDES) -o $@ $^

$(OBJDIR)/%.o: %.cpp
	@mkdir -p $(dir $@)
	@echo Compiling $<
	@$(CXX) $(CXXFLAGS) $(INCLUDES) -c $< -o $@

clean:
	@echo Cleaning objects
	@rm -rf $(OBJDIR)

fclean: clean
	@echo Cleaning $(PREDICT_NAME) $(TRAIN_NAME)
	@rm -rf $(PREDICT_NAME) $(TRAIN_NAME)

.PHONY: all clean fclean run re

-include $(TRAIN_DEPS) $(PREDICT_DEPS)
