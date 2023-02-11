
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:

Scenario: get request

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"files": [
			{
				"name": "image.png",
				"data": "iVBORw0KGgoAAAANSUhEUgAAABUAAAAUCAYAAABiS3YzAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JQAAgIMAAPn/AACA6QAAdTAAAOpgAAA6mAAAF2+SX8VGAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAC4jAAAuIwF4pT92AAAAB3RJTUUH5wILCBEHqfqy8AAABGhJREFUOMuVk11sU3UYxp///5yentOzrmXdurXd2NigjC/BweaGOlz8nEpQCSTEaDKThWggxgs/gkEv9UaJUYkSJUZjghglRkmM4wJDgrgxB4zxsbq1sK4fru3O2vWc09Nzzt8bQDBRxnP3XLy/PE/e9yW4Jvank4LY9Lm3fRWUI/LBL8pJAo3P52xu9+veUFFn09+9phgog9H1uo3/EW8PiwCAj78W+YYGi0vLnBlWTJMjmRUAqLsSs48v5x3Ook0RJBxLM8saEcHdrV9DEFTXtyFURZGbVTA1FQEPgAAgQ5NSzR8TcPzy5TQHIARgKQA3l4ewY6hgAphGNaZJN37HZwIPwPQEO7GybT1+O/qzM+BdFC4Z5QSALE94RkHBI6Yn95+Y3QDgsVyAGzl2vzwQy3HkifNFx6qUwcDDCx6+/QGvFBhgTA71zs9N/8ApTcvu27K9c5eWTWYL+eyrAMA/+aaH2R4OR39NBghFB/Pj+L513skrLkHiROBkg5iDg3J1tVb2k/7c6OBeKYRBWy3mVe7hbX0vh5sduyPndT6dntvscAgFAKCVbuIK23YrAzYyglGMS6eTQaGifYlldK4SnTUUwUrVElxZm9gZSvsPJmKDZ7xWd0/D+4IcfEsvoDoZyxw6OzpyKjJxGQBAJ2NG8aUfM+MEiBILo8Sj6dC0SLXXnOJ4zljdKEx1uJrj+/bvUqjIzAcW+YWaxvY9ra3uPokHV0gVM7Y587ngrDJvbP/CJZ1Jqq0BOP1sX8DHP898ojk3E0tVwGMhwWUXdcMQ5pbXHxgiKz3WuntW7Qx5pJ3UKBE30RBP5qPJdGrcKGVunBQtzGpYbDNQALzJLADWgcMa9rwzg0OHLRWaq6dWFr/317a9AfxF6kUW7vEocrOaQ6ujDM5plwr51C13y99sJIkq5CZ/IhZliwPej2rcrl5Jqur3NG098mL/o7KTmSimp1Axn8Js3jE/JsOeKvwH9NMDiX8/B/sm6lFqRV9J8NbVu6qEDzKkxrWxay3mkzFcHBkuzJwZ+PDqpac0Evr2n/q4jUSmlCwe45ahawKVe2JJu2NpcxPWdG2Cs345u5jTMzUPVTLZ17pwaNuO43pRjb6iZs9tK2UunByPxImul0AJQTBQ63a4WvZaRrELlDqvz3C3g1459S6ccmBedtZF8mp8zOR8z3R3tkh1fh8kSSRjlxPhybi9hRr6sJqfnFgQFACM+QQKc5dAeVayxNatPEf9m7rCcLlkwFYxPHiOZJLJr3TtamzB0OuySgXLXdfReyXNLRN5FS1N1QCzoEyfNRvLF46MK0KkbBRu3f5CQjNLG1IN2vvewfP46dgYJL6I5KTqyCS8a9XC2LFodMK4o6QA4JP9ishLm8vUU5mc5RHPCiiqIjVKdnvb+tUTT2958OIdQ/uQ4dfaqXZFKgctBxOooYDqCbSYcWe4Ubyroqlx4E7r44hRTmxfEXzhkQ1LOjXG36sW9TWU+kUhVR5xVFbGZZec+Rv1iPPNtWRaoAAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAyMy0wMi0xMFQxMDo1Mjo1OSswMDowMPMj0L4AAAAldEVYdGRhdGU6bW9kaWZ5ADIwMjMtMDItMTBUMTA6NTI6NTkrMDA6MDCCfmgCAAAAAElFTkSuQmCC"
			}
		],
		"commands": [
			{
			"command": "convert /tests/logo.png json:"
			},
			{
			"command": "convert /tests/logo.png -resize 100x100 new.png"
			},
			{
			"command": "ls -la"
			}
		],
		"return": [
			"new.png"
		]
	}
	"""
	When method POST
	Then status 200
	