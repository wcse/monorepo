
BASEDIR=$(pwd)
cd "${BASEDIR}"/wetalk || exit
yarn build

cd "$BASEDIR"/nextjs-dockerize || exit
yarn build
