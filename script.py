import pendulum

now = pendulum.now("Europe/Paris")

print(now.to_iso8601_string())