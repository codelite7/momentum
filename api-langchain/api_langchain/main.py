import os
from wsgiref.simple_server import make_server
from pyramid.config import Configurator
from api_langchain import openaii, health, anthropicc, groqq

if __name__ == '__main__':
    # .
    with Configurator() as config:
        config.add_route('openai', '/openai')
        config.add_view(openaii.openai, route_name='openai')
        config.add_route('groq', '/groq')
        config.add_view(groqq.groq, route_name='groq')
        config.add_route('anthropic', '/anthropic')
        config.add_view(anthropicc.anthropic, route_name='anthropic')
        config.add_route('health', '/health')
        config.add_view(health.health, route_name='health')
        app = config.make_wsgi_app()
    server = make_server('0.0.0.0', int(os.environ.get('PORT', '6543')), app)
    server.serve_forever()