import { NextSeo } from 'next-seo';
import Link from 'next/link';

import IdentityLayout from '../src/layouts/Identity';

const SignUp = () => {
  return (
    <>
      <NextSeo title="Wetalk | Sign Up" description="Talk together!" />

      <IdentityLayout>
        <div className="w-1/2 mx-auto flex flex-col p-10 bg-white rounded-lg shadow-xl">
          <div className="flex">
            <div className="w-1/2">
              <h1 className="text-3xl font-bold mb-1">REGISTER</h1>
              <p className="text-gray-700 mb-4">
                Lorem, ipsum dolor sit amet consectetur adipisicing elit. Quae
                cum similique exercitationem at nesciunt pariatur reprehenderit,
                consequuntur cumque libero architecto in aperiam aut assumenda
                nostrum doloribus nisi soluta quasi officia!
              </p>
              <div className="text-gray-500">
                <Link href="/signin">
                  <a className="text-blue-500 ml-2 font-bold">
                    Already have an account?
                  </a>
                </Link>
              </div>
            </div>

            <div className="w-1/2">
              <div className="flex flex-col justify-center h-full">
                <div className="w-full mb-4">
                  <button
                    className="w-full shadow bg-purple-500 hover:bg-purple-400 focus:outline-none text-white font-bold py-2 px-4 rounded"
                    type="submit"
                  >
                    Google
                  </button>
                </div>

                <div className="w-full">
                  <button
                    className="w-full shadow bg-blue-500 hover:bg-blue-400 focus:outline-none text-white font-bold py-2 px-4 rounded"
                    type="submit"
                  >
                    Facebook
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </IdentityLayout>
    </>
  );
};

export default SignUp;
