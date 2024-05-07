import { Injectable } from '@angular/core';
import { GraphQLClient } from 'graphql-request'

export interface listArgs {
  after?: string
  before?: string
  first?: number
  last?: number
  orderBy?: string[]
  where?: any
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

  getListReturnFields(returnFields: string): string {
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

  getListArgs(first?: number, last?: number, after?: string, before?: string, orderBy?: OrderBy[], where?: any): string {
    let argsString = ''
    if (first || last || after || before || orderBy || where) {
      let args: string[] = []
      if (first) {
        args.push(`first:${first}`)
      }
      if (last) {
        args.push(`last:${last}`)
      }
      if (after) {
        args.push(`after:"${after}"`)
      }
      if (before) {
        args.push(`before:"${before}"`)
      }
      if (orderBy) {
        args.push(`orderBy:${orderBy}`)
      }
      if (where) {
        args.push(`where:${where}`)
      }
      argsString = `(${args.join(',')})`
    }

    return argsString
  }
}
