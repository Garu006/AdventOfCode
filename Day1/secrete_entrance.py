import sys

def main():
    pos = 50    # posición inicial del dial
    count = 0   # cuántas veces termina en 0

    tokens = sys.stdin.read().split()

    for token in tokens:
        token = token.strip()
        if not token:
            continue  # saltar líneas vacías

        direction = token[0]      # 'L' o 'R'
        steps_str = token[1:]     # el número como texto

        try:
            steps = int(steps_str)
        except ValueError:
            print(f"Error: no pude convertir {steps_str!r} a entero (línea: {token!r})")
            return

        if direction == 'R':
            pos += steps
        elif direction == 'L':
            pos -= steps
        else:
            print(f"Error: dirección inválida {direction!r} en línea {token!r}")
            return

        # Normalizar a rango 0..99 (maneja negativos)
        pos = (pos % 100 + 100) % 100

        if pos == 0:
            count += 1

    print(count)


if __name__ == "__main__":
    main()
