import { inject as service } from '@ember/service';
import Route from 'consul-ui/routing/route';

export default class IndexRoute extends Route {
  @service('data-source/service') data;

  queryParams = {
    sortBy: 'sort',
    status: 'status',
    source: 'source',
    kind: 'kind',
    search: {
      as: 'filter',
      replace: true,
    },
  };

  async model(params, transition) {
    let terms = params.s || '';
    // we check for the old style `status` variable here
    // and convert it to the new style filter=status:critical
    let status = params.status;
    if (status) {
      status = `status:${status}`;
      if (terms.indexOf(status) === -1) {
        terms = terms
          .split('\n')
          .concat(status)
          .join('\n')
          .trim();
      }
    }
    terms = terms !== '' ? terms.split('\n') : [];
    const nspace = this.modelFor('nspace').nspace.substr(1);
    const dc = this.modelFor('dc').dc.Name;
    const items = await this.data.source(uri => uri`/${nspace}/${dc}/services`);
    return {
      dc,
      nspace,
      terms,
      items,
    };
  }

  setupController(controller, model) {
    super.setupController(...arguments);
    controller.setProperties(model);
  }
}
