import withLayout from 'hocs/withLayout';
import { NextSeo } from 'next-seo';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBoxOpen } from '@fortawesome/free-solid-svg-icons';
import Link from 'next/link';

import styles from './styles.module.scss';

const MyWall = () => {
  return (
    <>
      <NextSeo title="Wetalk | My Wall" description="Talk together!" />

      <div>
        <div
          className="h-64 bg-blue-700 bg-no-repeat bg-center bg-cover"
          style={{
            backgroundImage:
              'url(https://img.freepik.com/free-vector/neon-lights-wallpaper_52683-46462.jpg?size=626&ext=jpg)',
          }}
        ></div>
        <div className="flex flex-col items-center -mt-40 mb-8">
          <div
            className="w-40 h-40 bg-black rounded-full bg-no-repeat bg-center bg-cover border-4 border-solid border-black mb-1"
            style={{
              backgroundImage:
                'url(https://i1.sndcdn.com/avatars-000905704253-3tzsys-t200x200.jpg)',
            }}
          ></div>
          <div className="text-2xl font-bold">Huy Huynh</div>
        </div>
      </div>

      <nav className="flex items-center justify-between flex-wrap mb-4">
        <div className="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
          <div className="text-md lg:flex-grow h-tabs">
            <Link href="/mywall">
              <a className="block mt-4 lg:inline-block lg:mt-0 text-black mr-4 active">
                All
              </a>
            </Link>
            <Link href="/mywall/albums">
              <a className="block mt-4 lg:inline-block lg:mt-0 text-black mr-4">
                Albums
              </a>
            </Link>
            <Link href="/mywall/playlist">
              <a className="block mt-4 lg:inline-block lg:mt-0 text-black mr-4">
                Playlist
              </a>
            </Link>
          </div>
        </div>
      </nav>

      <div className="tab-content mb-5">
        <div
          className={`bg-white flex flex-col items-center justify-center ${styles.all_empty}`}
        >
          <FontAwesomeIcon
            icon={faBoxOpen}
            className="text-6xl mb-2 text-gray-500"
          />
          <div className="text-2xl mb-2">Seems a little quiet over here</div>
          <div className="text-sm mb-5">
            Upload a track to share it with your followers.
          </div>
          <Link href="/upload">
            <a className="shadow bg-blue-700 hover:bg-blue-600 focus:outline-none text-white font-bold py-2 px-4 rounded mb-10">
              Upload Now
            </a>
          </Link>
        </div>
      </div>
    </>
  );
};

export default withLayout(MyWall);
