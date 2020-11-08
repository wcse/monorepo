import Link from 'next/link';
import { useState } from 'react';

import styles from './styles.module.scss';

const Account = () => {
  const [show, setShow] = useState(false);

  const onToggle = () => {
    setShow((show) => !show);
  };

  const showHide = show ? 'block' : 'hidden';

  return (
    <div>
      <div className="flex items-center cursor-pointer" onClick={onToggle}>
        <div
          className="w-10 h-10 bg-black rounded-full mr-3 bg-no-repeat bg-center bg-cover"
          style={{
            backgroundImage:
              'url(https://i1.sndcdn.com/avatars-000905704253-3tzsys-t50x50.jpg)',
          }}
        ></div>
        <div>Huy Huynh</div>
        <svg
          className="-mr-1 ml-2 h-5 w-5"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fillRule="evenodd"
            d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
          />
        </svg>
      </div>

      <div
        className={`absolute bg-white w-56 text-black shadow-xl ${styles.sub_menu} ${showHide}`}
      >
        <ul>
          <li>
            <Link href="/mywall">
              <a>My Profile</a>
            </Link>
          </li>
          <li>
            <Link href="/signup">
              <a className={styles.logout}>Logout</a>
            </Link>
          </li>
        </ul>
      </div>

      <div
        className={`${styles.modal_backdrop} ${showHide}`}
        onClick={onToggle}
      ></div>
    </div>
  );
};

export default Account;
