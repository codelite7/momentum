import os

from langchain_core.messages import messages_from_dict
from pyramid.response import Response
from langchain_community.chat_models import ChatPerplexity
from langchain_core.prompts import ChatPromptTemplate

perplexityChat = ChatPerplexity(
    temperature=0,
    pplx_api_key=os.environ.get('PERPLEXITY_API_KEY', ''),
    model="llama-3-sonar-large-32k-online"
)
system = "You are a helpful assistant."


def perplexity(request):
    try:
        messages = messages_from_dict(request.json_body)
        response = perplexityChat.invoke(messages)
        content = response.content
        return Response(content)
    except Exception as e:
        print(e)
        raise e
