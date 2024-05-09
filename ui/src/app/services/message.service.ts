import { inject, Injectable } from '@angular/core'
import { gql } from 'graphql-request'
import { GraphqlService, OrderBy } from './graphql.service'

@Injectable({
  providedIn: 'root'
})
export class MessageService {
  graphqlService: GraphqlService = inject(GraphqlService)
  constructor() { }

  async createMessage(threadId: string, content: string): Promise<any>{
    let doc = gql`
    mutation createMessage($input:CreateMessageInput!) {
      createMessage(input:$input){
        id
          createdAt
          updatedAt
          content
          sentByUser {
            id
          }
          sentByAgent {
            id
          }
      }
    }
    `

    let response = await this.graphqlService.request(doc, {input: {threadID: threadId, content: content}})
    return response.createMessage
  }

  async getThreadMessages(threadId: string, first: number, orderBy?: OrderBy[], last?: number, after?: string, before?: string) {
    let doc = gql`
    query messages($after:Cursor, $first:Int,$before:Cursor,$last:Int, $threadId:ID!, $orderBy:[MessageOrder!]) {
      messages(after:$after, first:$first,before:$before, last:$last,where:{hasThreadWith:{id:$threadId}}, orderBy: $orderBy) {
        edges {
          node{
            id
            createdAt
            updatedAt
            content
            sentByUser {
              id
            }
            sentByAgent {
              id
            }
          }
        }
        pageInfo {
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
      }
    }
    `
    let vars = {
      threadId: threadId,
      first: first,
      last: last,
      after: after,
      before: before,
      orderBy: orderBy,
    }
    let response = await this.graphqlService.request(doc, vars)
    return response.messages
  }

  async listMessages(first: number, last?: number, after?: string, before?: string, orderBy?: OrderBy[], where?: any) {
    let doc = gql`
    query messages($after:Cursor, $first:Int,$before:Cursor,$last:Int,$where:MessageWhereInput) {
      messages(after:$after, first:$first,before:$before, last:$last,where:$where){
        edges {
          node {
            id
            createdAt
            updatedAt
            content
          }
        }
        pageInfo {
          hasNextPage
          hasPreviousPage
          startCursor
          endCursor
        }
      }
    }
    `
    let vars = {
      first: first,
      last: last,
      after: after,
      before: before,
      where: where,
    }
    return await this.graphqlService.request(doc, vars)
  }
}
