import { inject, Injectable } from '@angular/core'
import { GraphqlService, OrderBy } from './graphql.service'
import { gql } from 'graphql-request'

@Injectable({
  providedIn: 'root'
})
export class ThreadService {
  graphqlService: GraphqlService = inject(GraphqlService)

  constructor() {
  }

  async createThread(name: string) {
    let doc = gql`
    mutation createThread($input:CreateThreadInput!) {
      createThread(input:$input){
        id
        name
      }
    }
    `

    return await this.graphqlService.request(doc, {input: {name: name}})
  }

  async listThreads(first: number, returnFields: string, last?: number, after?: string, before?: string, orderBy?: OrderBy[], where?: any) {
    let args = this.graphqlService.getListArgs(first, last, after, before, orderBy, where)
    let fields = this.graphqlService.getListReturnFields(returnFields)
    let doc = gql`
    query threads {
        threads${args} ${fields}
      }
    `
    return await this.graphqlService.request(doc)
  }

  async getThread(id: string): Promise<any | undefined> {
    let doc = gql`
    query thread($id:ID!) {
      threads(where:{id:$id}) {
        edges {
          node {
            id
            createdAt
            updatedAt
            name
            parent {
              id
            }
            child {
              id
            }
          }
        }
      }
    }
    `
    let response = await this.graphqlService.request(doc, {id: id})
    if (response.threads?.edges?.length > 0) {
      return response.threads.edges[0].node
    }
    return undefined
  }

}
