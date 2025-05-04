package maps

var HP float // Значение уровня здоровья, от 0 до 1

func Fragment(_ vec4, texCoord vec2, _ vec4) vec4 {
    c := imageSrc0At(texCoord)    // Пиксель из спрайта здания
    mask := imageSrc1At(texCoord) // Пиксель из маски
    if c.a != 0.0 && mask.a != 0.0 {
        a := clamp(HP+(1.0-mask.a), 0.0, 1.0)
        // Создаём более тёмный пиксель при повреждениях.
        return vec4(c.r*a, c.g*a, c.b*a, c.a)
    }
    return c // Используем пиксель как есть
}