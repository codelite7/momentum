import os
from wsgiref.simple_server import make_server
from pyramid.config import Configurator
from api_langchain import perplexity, health

if __name__ == '__main__':
    # .
    with Configurator() as config:
        config.add_route('perplexity', '/perplexity')
        config.add_view(perplexity.perplexity, route_name='perplexity')
        config.add_route('health', '/health')
        config.add_view(health.health, route_name='health')
        app = config.make_wsgi_app()
    server = make_server('0.0.0.0', int(os.environ.get('PORT', '6543')), app)
    server.serve_forever()