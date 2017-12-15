def checkCaught(steps, depth):
    steps = steps + 1
    if (steps / depth) % 2 == 0:
        idx = depth - (-steps % depth) - 1
    else:
        idx = (depth - (steps % depth)) - 1
    return idx == 0

print checkCaught(0, 3)