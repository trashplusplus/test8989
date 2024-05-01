List of correct and incorrect inputs in /calculate endpoint

🍃 Correct POST -> http://localhost:8989/calculate
```
{
    "a": 5,
    "b": 7
}
```

🍃 Response 200: 
```
{
    "a": 120,
    "b": 5040
}
```

🍂 Incorrect POST -> http://localhost:8989/calculate

```
{
    "a": -4,
    "b": 2
}
```
🍂 Response 400
```
{
    "error": "Incorrect input"
}
```
🍂 Incorrect Post -> http://localhost:8989/calculate
```
{
    "a": 45,
    "b": 2
}
```
🍂 Response 400
```
{
    "error": "Factorial is too large"
}
```
🍂 Incorrect Post -> http://localhost:8989/calculate
```
{
    "x": 45,
    "y": 2
}
```
🍂 Reponse 400
```
{
    "error": "Incorrect input"
}
```



