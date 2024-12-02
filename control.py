def is_safe(row):
    inc = [row[i + 1] - row[i] for i in range(len(row) - 1)]
    if set(inc) <= {1, 2, 3} or set(inc) <= {-1, -2, -3}:
        return True
    return False

data = [[int(y) for y in x.split(' ')] for x in open('input.txt').read().split('\n')]

for row_idx, row in enumerate(data):
    for i in range(len(row)):
        test_row = row[:i] + row[i+1:]  # Remove one element
        if is_safe(test_row):
            print(f"Test Row {test_row}\n")
            break

safe_count = sum([any([is_safe(row[:i] + row[i + 1:]) for i in range(len(row))]) for row in data])
print(safe_count)