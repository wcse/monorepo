import { AppProps } from 'next/app';
import { Provider } from 'react-redux';
import store from '../src/redux/store';

import '../styles/tailwind.scss';
import '../styles/globals.css';

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <Provider store={store}>
      <Component {...pageProps} />
    </Provider>
  );
}

export default MyApp;
