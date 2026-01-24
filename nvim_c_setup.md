# Neovim C Development Setup Guide

## Quick Start: Installing Required Tools

```bash
# Compiler and build tools (if not already installed)
sudo apt install build-essential gdb valgrind clang-format

# Install clangd (C/C++ LSP server)
sudo apt install clangd

# For Neovim plugin management, if you don't have a plugin manager:
# Using lazy.nvim (modern, popular choice)
# Using packer.nvim (also good)
# Using vim-plug (simpler)
```

## Neovim Configuration for C Development

Add this to your `~/.config/nvim/init.lua` (or adapt for your existing setup):

```lua
-- Lazy.nvim plugin manager setup (if using lazy.nvim)
-- If you already have a plugin manager, adapt accordingly

-- Basic LSP setup
require('lspconfig').clangd.setup({
    cmd = { "clangd", "--background-index" },
    capabilities = capabilities,
    on_attach = function(client, bufnr)
        -- Enable completion triggered by <c-x><c-o>
        vim.api.nvim_buf_set_option(bufnr, 'omnifunc', 'v:lua.vim.lsp.omnifunc')
        
        -- Key mappings for LSP
        local bufopts = { noremap=true, silent=true, buffer=bufnr }
        vim.keymap.set('n', 'gd', vim.lsp.buf.definition, bufopts)
        vim.keymap.set('n', 'K', vim.lsp.buf.hover, bufopts)
        vim.keymap.set('n', 'gi', vim.lsp.buf.implementation, bufopts)
        vim.keymap.set('n', 'gr', vim.lsp.buf.references, bufopts)
        vim.keymap.set('n', '<leader>rn', vim.lsp.buf.rename, bufopts)
        vim.keymap.set('n', '<leader>ca', vim.lsp.buf.code_action, bufopts)
        vim.keymap.set('n', '<leader>f', function() vim.lsp.buf.format { async = true } end, bufopts)
    end,
})

-- Auto-format on save
vim.api.nvim_create_autocmd("BufWritePre", {
    pattern = {"*.c", "*.h"},
    callback = function()
        vim.lsp.buf.format({ async = false })
    end,
})

-- Recommended C-specific settings
vim.api.nvim_create_autocmd("FileType", {
    pattern = "c",
    callback = function()
        vim.opt_local.tabstop = 4
        vim.opt_local.shiftwidth = 4
        vim.opt_local.expandtab = true
        vim.opt_local.cindent = true
    end,
})
```

### Essential Neovim Plugins for C Development

Add these to your plugin manager configuration:

```lua
-- For lazy.nvim:
{
    'neovim/nvim-lspconfig',
    dependencies = {
        'hrsh7th/nvim-cmp',         -- Autocompletion
        'hrsh7th/cmp-nvim-lsp',     -- LSP source for nvim-cmp
        'hrsh7th/cmp-buffer',       -- Buffer completions
        'L3MON4D3/LuaSnip',         -- Snippet engine
    }
},
{
    'nvim-telescope/telescope.nvim',  -- Fuzzy finder
    dependencies = { 'nvim-lua/plenary.nvim' }
},
{
    'nvim-treesitter/nvim-treesitter',  -- Better syntax highlighting
    build = ':TSUpdate'
},
```

### Minimal .clang-format Configuration

Create `~/.clang-format` or put in your project root:

```yaml
BasedOnStyle: LLVM
IndentWidth: 4
ColumnLimit: 100
BreakBeforeBraces: Linux
AllowShortFunctionsOnASingleLine: Empty
IndentCaseLabels: false
PointerAlignment: Right
```

## Project Structure for DSA Learning

```
dsa-c/
├── Makefile
├── .clang-format
├── src/
│   ├── main.c
│   └── utils/
│       ├── input.c
│       └── input.h
├── data_structures/
│   ├── linked_list.c
│   ├── linked_list.h
│   ├── binary_tree.c
│   └── binary_tree.h
├── algorithms/
│   ├── sorting.c
│   └── sorting.h
├── tests/
│   └── test_*.c
└── build/
    └── (compiled objects)
```

## Makefile Template

```makefile
CC = gcc
CFLAGS = -Wall -Wextra -Wpedantic -std=c11 -Idata_structures -Ialgorithms -Isrc
DEBUG_FLAGS = -g -O0 -fsanitize=address,undefined
RELEASE_FLAGS = -O2 -DNDEBUG

SRCDIR = src
DSDIR = data_structures
ALGODIR = algorithms
BUILDDIR = build

SRCS = $(wildcard $(SRCDIR)/*.c) $(wildcard $(SRCDIR)/*/*.c)
DS_SRCS = $(wildcard $(DSDIR)/*.c)
ALGO_SRCS = $(wildcard $(ALGODIR)/*.c)

ALL_SRCS = $(SRCS) $(DS_SRCS) $(ALGO_SRCS)
OBJS = $(ALL_SRCS:%.c=$(BUILDDIR)/%.o)

TARGET = dsa_program

# Default: debug build
all: CFLAGS += $(DEBUG_FLAGS)
all: directories $(TARGET)

release: CFLAGS += $(RELEASE_FLAGS)
release: directories $(TARGET)

directories:
	@mkdir -p $(BUILDDIR)/$(SRCDIR) $(BUILDDIR)/$(DSDIR) $(BUILDDIR)/$(ALGODIR)

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) $^ -o $@

$(BUILDDIR)/%.o: %.c
	@mkdir -p $(dir $@)
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -rf $(BUILDDIR) $(TARGET)

run: all
	./$(TARGET)

debug: all
	gdb ./$(TARGET)

valgrind: all
	valgrind --leak-check=full --show-leak-kinds=all ./$(TARGET)

.PHONY: all release clean run debug valgrind directories
```

## Basic C Input/Output Commands

### Output Functions (from stdio.h)

```c
#include <stdio.h>

// Basic output
printf("Hello, World!\n");                    // Print with newline
printf("Number: %d\n", 42);                   // Integer
printf("Float: %.2f\n", 3.14159);            // Float with 2 decimals
printf("String: %s\n", "Hello");              // String
printf("Char: %c\n", 'A');                    // Single character
printf("Pointer: %p\n", (void*)&variable);    // Memory address

// Multiple values
printf("x=%d, y=%d\n", x, y);

// Output without newline
printf("Enter name: ");  // Cursor stays on same line

// Error output
fprintf(stderr, "Error: File not found\n");

// Formatted output to string
char buffer[100];
sprintf(buffer, "Value: %d", 42);
snprintf(buffer, sizeof(buffer), "Safe: %d", 42);  // Safer (size-limited)
```

### Input Functions (from stdio.h)

```c
#include <stdio.h>

// Reading integers
int num;
printf("Enter a number: ");
scanf("%d", &num);  // Note the & (address-of operator)

// Reading multiple integers
int a, b;
scanf("%d %d", &a, &b);  // Reads space-separated integers

// Reading floats/doubles
float f;
double d;
scanf("%f", &f);   // float
scanf("%lf", &d);  // double (note: lf for scanf, f for printf)

// Reading characters
char c;
scanf(" %c", &c);  // Space before %c skips whitespace

// Reading strings (DANGEROUS - no bounds check)
char name[50];
scanf("%s", name);  // Stops at whitespace, NO & needed for arrays

// Reading strings safely (recommended)
scanf("%49s", name);  // Reads at most 49 chars + null terminator

// Reading entire line (including spaces)
char line[100];
fgets(line, sizeof(line), stdin);  // Safer than gets()
// Note: fgets includes newline, remove it:
line[strcspn(line, "\n")] = '\0';

// Clear input buffer after scanf (common issue)
int ch;
while ((ch = getchar()) != '\n' && ch != EOF);
```

### Common scanf Gotchas and Solutions

```c
// Problem: scanf leaves newline in buffer
int num;
scanf("%d", &num);      // User types "5\n"
char c = getchar();     // Gets '\n' instead of next input!

// Solution 1: Clear buffer
int num;
scanf("%d", &num);
while (getchar() != '\n');  // Clear buffer
char c = getchar();

// Solution 2: Use space before %c
scanf("%d", &num);
scanf(" %c", &c);  // Space skips whitespace including '\n'

// Problem: scanf fails if input doesn't match format
if (scanf("%d", &num) != 1) {
    fprintf(stderr, "Invalid input\n");
    while (getchar() != '\n');  // Clear bad input
}
```

## Complete Starter Program

Here's a working program demonstrating all I/O basics:

```c
/* src/main.c */
#include <stdio.h>
#include <string.h>

void demo_output(void) {
    printf("=== Output Examples ===\n");
    
    int age = 25;
    float height = 5.9;
    char initial = 'J';
    char name[] = "John";
    
    printf("Integer: %d\n", age);
    printf("Float: %.1f\n", height);
    printf("Character: %c\n", initial);
    printf("String: %s\n", name);
    printf("Multiple: Name=%s, Age=%d\n", name, age);
    printf("\n");
}

void demo_input_integers(void) {
    printf("=== Integer Input ===\n");
    
    int x, y;
    printf("Enter two integers (space-separated): ");
    
    if (scanf("%d %d", &x, &y) != 2) {
        fprintf(stderr, "Error: Invalid input\n");
        while (getchar() != '\n');  // Clear buffer
        return;
    }
    
    printf("You entered: %d and %d\n", x, y);
    printf("Sum: %d\n\n", x + y);
    
    while (getchar() != '\n');  // Clear buffer for next input
}

void demo_input_string(void) {
    printf("=== String Input ===\n");
    
    char name[50];
    printf("Enter your name: ");
    
    if (fgets(name, sizeof(name), stdin) != NULL) {
        // Remove trailing newline
        name[strcspn(name, "\n")] = '\0';
        printf("Hello, %s!\n\n", name);
    }
}

void demo_menu_system(void) {
    printf("=== Simple Menu ===\n");
    
    int choice;
    while (1) {
        printf("\n1. Print message\n");
        printf("2. Do calculation\n");
        printf("3. Exit\n");
        printf("Choose (1-3): ");
        
        if (scanf("%d", &choice) != 1) {
            fprintf(stderr, "Invalid input\n");
            while (getchar() != '\n');
            continue;
        }
        while (getchar() != '\n');  // Clear buffer
        
        switch (choice) {
            case 1:
                printf("Hello from option 1!\n");
                break;
            case 2:
                printf("2 + 2 = %d\n", 2 + 2);
                break;
            case 3:
                printf("Exiting...\n");
                return;
            default:
                printf("Invalid choice\n");
        }
    }
}

int main(void) {
    demo_output();
    demo_input_integers();
    demo_input_string();
    demo_menu_system();
    
    return 0;
}
```

## Quick Start Commands

```bash
# Create project structure
mkdir -p dsa-c/{src,data_structures,algorithms,tests,build}
cd dsa-c

# Create files (use nvim)
nvim src/main.c
nvim Makefile

# Compile and run
make run

# Debug with GDB
make debug

# Check for memory leaks
make valgrind

# Clean build artifacts
make clean
```

## Essential Neovim Commands While Coding

```
# LSP Commands (after hovering over code)
gd          - Go to definition
K           - Show documentation/hover info
gr          - Show references
<leader>rn  - Rename symbol
<leader>ca  - Code actions
<leader>f   - Format file

# Building from Neovim
:!make            - Run make
:!make run        - Build and run
:make             - Use Neovim's built-in make integration

# Terminal in Neovim
:term             - Open terminal in split
:vertical term    - Open terminal vertically
<C-\><C-n>        - Exit terminal mode (back to normal mode)
```

## Your First DSA Program: Interactive Linked List

Create this to test your setup:

```c
/* data_structures/linked_list.h */
#ifndef LINKED_LIST_H
#define LINKED_LIST_H

typedef struct Node {
    int data;
    struct Node *next;
} Node;

Node *node_create(int value);
void list_print(Node *head);
void list_free(Node *head);
void list_append(Node **head, int value);

#endif

/* data_structures/linked_list.c */
#include <stdio.h>
#include <stdlib.h>
#include "linked_list.h"

Node *node_create(int value) {
    Node *node = malloc(sizeof(Node));
    if (node) {
        node->data = value;
        node->next = NULL;
    }
    return node;
}

void list_print(Node *head) {
    if (!head) {
        printf("List is empty\n");
        return;
    }
    
    Node *current = head;
    while (current) {
        printf("%d", current->data);
        if (current->next) printf(" -> ");
        current = current->next;
    }
    printf("\n");
}

void list_free(Node *head) {
    while (head) {
        Node *next = head->next;
        free(head);
        head = next;
    }
}

void list_append(Node **head, int value) {
    Node *new_node = node_create(value);
    if (!new_node) return;
    
    if (*head == NULL) {
        *head = new_node;
        return;
    }
    
    Node *current = *head;
    while (current->next) {
        current = current->next;
    }
    current->next = new_node;
}

/* src/main.c */
#include <stdio.h>
#include <stdlib.h>
#include "linked_list.h"

int main(void) {
    Node *head = NULL;
    int choice, value;
    
    printf("Linked List Interactive Demo\n");
    printf("============================\n\n");
    
    while (1) {
        printf("\n1. Add element\n");
        printf("2. Print list\n");
        printf("3. Exit\n");
        printf("Enter choice: ");
        
        if (scanf("%d", &choice) != 1) {
            fprintf(stderr, "Invalid input\n");
            while (getchar() != '\n');
            continue;
        }
        while (getchar() != '\n');
        
        switch (choice) {
            case 1:
                printf("Enter value: ");
                if (scanf("%d", &value) == 1) {
                    list_append(&head, value);
                    printf("Added %d to list\n", value);
                }
                while (getchar() != '\n');
                break;
                
            case 2:
                printf("Current list: ");
                list_print(head);
                break;
                
            case 3:
                printf("Cleaning up and exiting...\n");
                list_free(head);
                return 0;
                
            default:
                printf("Invalid choice\n");
        }
    }
    
    return 0;
}
```

## Testing Your Setup

```bash
# 1. Create the project structure
mkdir -p dsa-c/{src,data_structures,build}
cd dsa-c

# 2. Copy the files above into their locations

# 3. Compile and run
make run

# 4. Test the interactive program
# Try: Add elements (1), print list (2), exit (3)

# 5. Check for memory leaks
make valgrind
# Should show: "All heap blocks were freed -- no leaks are possible"
```

You're now ready to implement data structures and algorithms in C with a professional development environment.
