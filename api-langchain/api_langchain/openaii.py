import os

from langchain_core.messages import messages_from_dict
from langchain_openai import ChatOpenAI
from pyramid.response import Response

chat = ChatOpenAI(
    temperature=0,
    model="gpt-4o"
)
system = "You are a helpful assistant."


def openai(request):
    try:
        messages = messages_from_dict(request.json_body)
        response = chat.invoke(messages)
        content = response.content
        return Response(content)
    except Exception as e:
        print(e)
        raise e
