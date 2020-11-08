import { Component } from 'react';

import AdminLayout from 'layouts/admin';

const withLayout = (WrapperComponent, LayoutComponent = AdminLayout) =>
  class extends Component {
    render() {
      return (
        <LayoutComponent>
          <WrapperComponent />
        </LayoutComponent>
      );
    }
  };

export default withLayout;
