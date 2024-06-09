import os

from langchain_core.messages import messages_from_dict
from langchain_groq import ChatGroq
from pyramid.response import Response

chat = ChatGroq(
    temperature=0,
    model="llama3-8b-8192"
)
system = "You are a helpful assistant."


def groq(request):
    try:
        messages = messages_from_dict(request.json_body)
        response = chat.invoke(messages)
        content = response.content
        return Response(content)
    except Exception as e:
        print(e)
        raise e
