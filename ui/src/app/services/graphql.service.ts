import { Injectable } from '@angular/core';
import { GraphQLClient } from 'graphql-request'
import { getSdk } from '../../../graphql/generated'

export interface listArgs {
  after?: string
  before?: string
  first?: number
  last?: number
  orderBy?: string[]
  where?: any
}

export interface ListArgs {
  after?: string
  first?: number
  before?: string
  last?: number
  orderBy?: OrderBy[]
  where?: {}
}

export enum OrderByDirection {
  Ascending = 'ASC',
  Descending = 'DESC',
}

export interface OrderBy {
  direction: OrderByDirection
  field: string
}

@Injectable({
  providedIn: 'root'
})
export class GraphqlService {
  client: GraphQLClient = new GraphQLClient('http://localhost:3000/query');
  sdk = getSdk(this.client)
  constructor() { }

  async request(document: any, variables?: any): Promise<any> {
    let response: any
    if (variables) {
      response = await this.client.request(document, variables)
    } else {
      response = await(this.client.request(document))
    }
    return response
  }

  getQueryReturnFields(returnFields: string): string {
    return `
    {
      edges {
        node ${returnFields}
        cursor
      }
      pageInfo {
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
      totalCount
    }`
  }
}
