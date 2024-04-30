script_dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd ) 
echo $script_dir
infra_dir=$(dirname -- "$script_dir") 
echo $infra_dir
root_dir=$(dirname -- "$infra_dir") 
echo $root_dir
DOCKER_FILE_PATH=$script_dir/Dockerfile
PREFIX=beclab

docker  build    \
    -f ${DOCKER_FILE_PATH} \
    -t ${PREFIX}/article_extractor_develop $root_dir

# docker run --name article_extractor_develop -v /home/ubuntu/article-extractor:/opt/article-extractor --net=host -d beclab/article_extractor_develop 