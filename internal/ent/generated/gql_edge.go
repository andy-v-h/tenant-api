// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (t *Tenant) Parent(ctx context.Context) (*Tenant, error) {
	result, err := t.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Tenant) Children(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *TenantOrder, where *TenantWhereInput,
) (*TenantConnection, error) {
	opts := []TenantPaginateOption{
		WithTenantOrder(orderBy),
		WithTenantFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := t.Edges.totalCount[1][alias]
	if nodes, err := t.NamedChildren(alias); err == nil || hasTotalCount {
		pager, err := newTenantPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &TenantConnection{Edges: []*TenantEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return t.QueryChildren().Paginate(ctx, after, first, before, last, opts...)
}
